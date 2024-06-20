package model

type Tweet struct {
	Id        string `json:"id"`
	Uid       string `json:"uid"`
	Content   string `json:"content"`
	Date      string `json:"date"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}

type Reply struct {
	Id        string `json:"id"`
	TweetId   string `json:"tweet_id"`
	Uid       string `json:"uid"`
	Content   string `json:"content"`
	Date      string `json:"date"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}

type ReplyToReply struct {
	Id        int    `json:"id"`
	ReplyId   int    `json:"replyId"`
	Uid       string `json:"uid"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	LikeCount int    `json:"likeCount"`
}

type UpdateTweet struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}
