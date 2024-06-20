package model

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

type ReplyLike struct {
	ReplyId int    `json:"replyId"`
	Uid     string `json:"uid"`
}

type LikedReply struct {
	ID        int      `json:"id"`
	LikeCount int      `json:"likeCount"`
	Likes     []string `json:"likes"`
}
