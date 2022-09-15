package collections

const (
	ChallengeSlug string = "challenge"
	EntrySlug     string = "entry"
	UsernameSlug  string = "username"
)

type EntryCollectionTemplate struct {
	Description     string `json:"description"`
	CreatedByUserId string `json:"createdByUserId"`
	ChallengeId     string `json:"challengeId"`
}

type ChallengeCollectionTemplate struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	CreatedByUserId string `json:"createdByUserId"`
}

type UsernameCollectionTemplate struct {
	Username string `json:"username"`
}
