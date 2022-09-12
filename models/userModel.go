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
