package models

import (
	"errors"
	"github.com/GetStream/stream-go2/v7"
)

type EntryModel struct {
	ID             string         `json:"id"`
	Description    string         `json:"description"`
	CreatedBy      UserModel      `json:"createdBy"`
	ChallengeModel ChallengeModel `json:"challenge"`
}

func EntryModelFromCollectionsAndActor(entryCollection *stream.CollectionObjectResponse,
	actor *stream.UserResponse, challengeCollection *stream.CollectionObjectResponse) EntryModel {
	return EntryModel{
		ID:             entryCollection.ID,
		Description:    entryCollection.Data["description"].(string),
		CreatedBy:      UserModelFromGetStreamActor(actor),
		ChallengeModel: ChallengeModelFromCollectionAndActor(challengeCollection, actor),
	}
}

func EntryModelFromEnrichedActivity(activity stream.EnrichedActivity) (EntryModel, error) {

	extra := activity.Extra["challengeCollection"]

	if extra == nil {
		return EntryModel{}, errors.New("challengeCollection not found in activity.Extra")
	}

	challengeCollection := extra.(map[string]interface{})
	actor := activity.Actor

	//need to convert to stream.Data for challengeModelFromData to work
	challengeData := stream.Data{
		ID:    challengeCollection["id"].(string),
		Extra: challengeCollection,
	}

	challengeModel, err := ChallengeModelFromData(challengeData, actor)

	if err != nil {
		return EntryModel{}, err
	}

	return EntryModel{
		ID:             activity.Object.ID,
		Description:    activity.Object.Extra["data"].(map[string]interface{})["description"].(string),
		CreatedBy:      UserModelFromData(actor),
		ChallengeModel: challengeModel,
	}, nil
}
