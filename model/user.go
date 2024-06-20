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

type AvatarRequest struct {
	UserId    string `json:"userId"`
	AvatarURL string `json:"avatarURL"`
}

type UserProfile struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}
