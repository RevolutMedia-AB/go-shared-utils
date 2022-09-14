package models

import "github.com/GetStream/stream-go2/v7"

type EntryModel struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedBy   UserModel `json:"createdBy"`
	ChallengeModel ChallengeModel `json:"challenge"`
}


func EntryModelFromCollectionsAndActor(entryCollection *stream.CollectionObjectResponse, actor *stream.UserResponse, challengeCollection *stream.CollectionObjectResponse) EntryModel {
	return EntryModel{
		ID:				entryCollection.ID,
		Description:	entryCollection.Data["description"].(string),
		CreatedBy:		UserModelFromGetStreamActor(actor),
		ChallengeModel: ChallengeModelFromCollectionAndActor(challengeCollection, actor),
	}
}
