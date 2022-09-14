package models

import "github.com/GetStream/stream-go2/v7"

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

func ChallengeModelFromActivityExtra(collection map[string]interface{}, actor stream.Data) ChallengeModel {
	challengeCollection := collection["data"].(map[string]interface{})

	return ChallengeModel{
		ID:          challengeCollection["id"].(string),
		Title:       challengeCollection["title"].(string),
		Description: challengeCollection["description"].(string),
		CreatedBy:   UserModelFromActivityActor(actor),
	}
}
