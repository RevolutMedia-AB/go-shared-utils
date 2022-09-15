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

type FeedGroup struct {
	Feed    *stream.FlatFeed
	FeedRef string
}

func NewFeedGroup(client *stream.Client, slug string, id string) *FeedGroup {
	feed, err := client.FlatFeed(slug, id)
	if err != nil {
		panic(err)
	}
	return &FeedGroup{
		Feed:    feed,
		FeedRef: feed.Slug() + ":" + feed.UserID(),
	}
}

func ChallengeRaw(client *stream.Client, challengeId string) *FeedGroup {
	return NewFeedGroup(client, ChallengeRawSlug, challengeId)
}

func ChallengeNotification(client *stream.Client, challengeId string) *FeedGroup {
	return NewFeedGroup(client, ChallengeNotificationSlug, challengeId)
}

func ChallengeEntry(client *stream.Client, challengeId string) *FeedGroup {
	return NewFeedGroup(client, ChallengeEntrySlug, challengeId)
}

func GlobalChallenge(client *stream.Client) *FeedGroup {
	return NewFeedGroup(client, GlobalSlug, "challenge")
}

func UserRaw(client *stream.Client, userId string) *FeedGroup {
	return NewFeedGroup(client, UserRawSlug, userId)
}

func UserNotification(client *stream.Client, userId string) *FeedGroup {
	return NewFeedGroup(client, UserNotificationSlug, userId)
}

func UserChallenge(client *stream.Client, userId string) *FeedGroup {
	return NewFeedGroup(client, UserChallengeSlug, userId)
}

func UserChallengeEntry(client *stream.Client, userId string, challengeId string) *FeedGroup {
	return NewFeedGroup(client, UserChallengeEntrySlug, userId+"_"+challengeId)
}

func UserEntry(client *stream.Client, userId string) *FeedGroup {
	return NewFeedGroup(client, UserEntrySlug, userId)
}

func ProjectedUserHome(client *stream.Client, userId string) *FeedGroup {
	return NewFeedGroup(client, ProjectedUserHomeSlug, userId)
}

func ProjectedUserNotification(client *stream.Client, userId string) *FeedGroup {
	return NewFeedGroup(client, ProjectedUserNotificationSlug, userId)
}

func Curated(client *stream.Client, userId string) *FeedGroup {
	return NewFeedGroup(client, CuratedSlug, userId)
}
