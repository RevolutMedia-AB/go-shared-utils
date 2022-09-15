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

	challengeCollection := activity.Extra["challengeCollection"]
	challengeCreatedBy := activity.Extra["challengeCreatedBy"]

	if challengeCollection == nil {
		return EntryModel{}, errors.New("challengeCollection not found in activity.Extra")
	}

	if challengeCreatedBy == nil {
		return EntryModel{}, errors.New("challengeCreatedBy not found in activity.Extra")
	}

	challengeCollectionData := challengeCollection.(map[string]interface{})
	challengeCreatedByData := challengeCreatedBy.(map[string]interface{})

	challengeActorData := stream.Data{
		ID:    challengeCreatedByData["id"].(string),
		Extra: challengeCreatedByData,
	}

	//need to convert to stream.Data for challengeModelFromData to work
	challengeData := stream.Data{
		ID:    challengeCollectionData["id"].(string),
		Extra: challengeCollectionData,
	}

	challengeModel, err := ChallengeModelFromData(challengeData, challengeActorData) //user who created the challenge 

	if err != nil {
		return EntryModel{}, err
	}

	return EntryModel{
		ID:             activity.Object.ID,
		Description:    activity.Object.Extra["data"].(map[string]interface{})["description"].(string),
		CreatedBy:      UserModelFromData(activity.Actor),
		ChallengeModel: challengeModel,
	}, nil
}
