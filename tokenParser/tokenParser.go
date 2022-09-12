package tokenParserMiddleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type tokenData struct {
	AuthUserId string `json:"user_id"`
}

func TokenParser(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		authToken := r.Header.Get("Authorization")
		gatewayToken := r.Header.Get("x-apigateway-api-userinfo")
		if authToken == "" && gatewayToken == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		authUserId, err := getUserIdFromTokens(authToken, gatewayToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		newParam := httprouter.Param{Key: "authUserId", Value: authUserId}
		ps = append(ps, newParam)
		next(w, r, ps)
	}
}

func getUserIdFromTokens(authToken string, gatewayToken string) (string, error) {
	AuthUserId := ""
	err := error(nil)
	if authToken != "" {
		AuthUserId, err = decodeAuthToken(authToken)
	} else if gatewayToken != "" {
		AuthUserId, err = decodeGatewayToken(gatewayToken)
	}
	return AuthUserId, err
}

func decodeAuthToken(token string) (string, error) {
	cleanToken := strings.Replace(token, "Bearer ", "", 1)
	splitToken := strings.Split(cleanToken, ".")
	data, err := base64.RawURLEncoding.DecodeString(splitToken[1])
	if err != nil {
		return "", err
	}
	tokenDataRef := tokenData{}
	if err = json.NewDecoder(strings.NewReader(string(data))).Decode(&tokenDataRef); err != nil {
		return "", err
	}
	return tokenDataRef.AuthUserId, nil
}

func decodeGatewayToken(token string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	tokenDataRef := tokenData{}
	if err = json.NewDecoder(strings.NewReader(string(data))).Decode(&tokenDataRef); err != nil {
		return "", err
	}
	return tokenDataRef.AuthUserId, nil
}
