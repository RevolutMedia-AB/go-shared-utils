package collections

const (
	FollowUserSlug      string = "follow_user"
	FollowChallengeSlug string = "follow_challenge"
	ChallengeSlug       string = "challenge"
	EntrySlug           string = "entry"
	UsernameSlug        string = "username"
	CommentSlug         string = "comment"
)

type Collections struct {
	FollowUser      string `default:"follow_user"`
	FollowChallenge string `default:"follow_challenge"`
	Challenge       string `default:"challenge"`
	Entry           string `default:"entry"`
	Username        string `default:"username"`
}
