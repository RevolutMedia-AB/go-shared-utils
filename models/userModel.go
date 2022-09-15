package models

import (
	"github.com/GetStream/stream-go2/v7"
)

type UserModel struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
}

func (u *UserModel) ToJSON() {

}

func UserModelFromGetStreamActor(user *stream.UserResponse) UserModel {
	return UserModel{
		ID:                user.ID,
		Username:          user.Data["username"].(string),
		ProfilePictureUrl: user.Data["profilePictureUrl"].(string),
	}
}

func UserModelFromData(actorData stream.Data) UserModel {
	data := actorData.Extra["data"].(map[string]interface{})
	return UserModel{
		ID:       actorData.ID,
		Username: data["username"].(string),
		ProfilePictureUrl: data["profilePictureUrl"].(string),
	}
}
