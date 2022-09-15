package models

import (
	"errors"
	"github.com/GetStream/stream-go2/v7"
)

type ChallengeModel struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   UserModel `json:"createdBy"`
}

func ChallengeModelFromCollectionAndActor(collection *stream.CollectionObjectResponse, actor *stream.UserResponse) ChallengeModel {
	return ChallengeModel{
		ID:          collection.ID,
		Title:       collection.Data["title"].(string),
		Description: collection.Data["description"].(string),
		CreatedBy:   UserModelFromGetStreamActor(actor),
	}
}

func ChallengeModelFromData(collectionData stream.Data, actorData stream.Data) (ChallengeModel, error) {

	if extra := collectionData.Extra["data"]; extra != nil {
		data := extra.(map[string]interface{})

		return ChallengeModel{
			ID:          collectionData.ID,
			Title:       data["title"].(string),
			Description: data["description"].(string),
			CreatedBy:   UserModelFromData(actorData),
		}, nil
	}

	return ChallengeModel{}, errors.New("no data in challengeCollection")
}
