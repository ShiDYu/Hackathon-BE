package model

type UserGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserRegister struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`
}

type Tweet struct {
	Id       string `json:"id"`
	Uid      string `json:"uid"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Nickname string `json:"nickname"`
}

type Reply struct {
	Id      string `json:"id"`
	TweetId string `json:"tweet_id"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type Like struct {
	PostID int    `json:"postId"`
	UserID string `json:"userId"`
}

type LikedPost struct {
	ID        int      `json:"id"`
	LikeCount int      `json:"likeCount"`
	Likes     []string `json:"likes"`
}

type PostResponse struct {
	LikeCount   int  `json:"likeCount"`
	LikedByUser bool `json:"likedByUser"`
}
