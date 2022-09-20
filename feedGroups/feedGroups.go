package feedGroups

import "github.com/GetStream/stream-go2/v7"

const (
	ChallengeRawSlug              string = "challenge_raw"
	UserRawSlug                   string = "user_raw"
	UserNotificationSlug          string = "user_notification"
	ChallengeNotificationSlug     string = "challenge_notification"
	ChallengeEntrySlug            string = "challenge_entry"
	UserChallengeEntrySlug        string = "user_challenge_entry"
	UserChallengeSlug             string = "user_challenge"
	ProjectedUserNotificationSlug string = "projected_user_notification"
	ProjectedUserHomeSlug         string = "projected_user_home"
	UserEntrySlug                 string = "user_entry"
	GlobalSlug                    string = "global"
	CuratedSlug                   string = "curated"
)

type FlatFeedGroup struct {
	Feed    *stream.FlatFeed
	FeedRef string
}

type NotificationFeedGroup struct {
	Feed    *stream.NotificationFeed
	FeedRef string
}

func NewFlatFeedGroup(client *stream.Client, slug string, id string) *FlatFeedGroup {
	feed, err := client.FlatFeed(slug, id)
	if err != nil {
		panic(err)
	}
	return &FlatFeedGroup{
		Feed:    feed,
		FeedRef: feed.Slug() + ":" + feed.UserID(),
	}
}

func NewNotificationFeedGroup(client *stream.Client, slug string, id string) *NotificationFeedGroup {
	feed, err := client.NotificationFeed(slug, id)
	if err != nil {
		panic(err)
	}
	return &NotificationFeedGroup{
		Feed:    feed,
		FeedRef: feed.Slug() + ":" + feed.UserID(),
	}
}

func ChallengeRaw(client *stream.Client, challengeId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, ChallengeRawSlug, challengeId)
}

func ChallengeNotification(client *stream.Client, challengeId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, ChallengeNotificationSlug, challengeId)
}

func ChallengeEntry(client *stream.Client, challengeId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, ChallengeEntrySlug, challengeId)
}

func GlobalChallenge(client *stream.Client) *FlatFeedGroup {
	return NewFlatFeedGroup(client, GlobalSlug, "challenge")
}

func UserRaw(client *stream.Client, userId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, UserRawSlug, userId)
}

func UserNotification(client *stream.Client, userId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, UserNotificationSlug, userId)
}

func UserChallenge(client *stream.Client, userId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, UserChallengeSlug, userId)
}

func UserChallengeEntry(client *stream.Client, userId string, challengeId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, UserChallengeEntrySlug, userId+"_"+challengeId)
}

func UserEntry(client *stream.Client, userId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, UserEntrySlug, userId)
}

func ProjectedUserHome(client *stream.Client, userId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, ProjectedUserHomeSlug, userId)
}

func ProjectedUserNotification(client *stream.Client, userId string) *NotificationFeedGroup {
	return NewNotificationFeedGroup(client, ProjectedUserNotificationSlug, userId)
}

func Curated(client *stream.Client, userId string) *FlatFeedGroup {
	return NewFlatFeedGroup(client, CuratedSlug, userId)
}
