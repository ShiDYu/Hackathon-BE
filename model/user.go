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
	Id      string `json:"id"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
	Date    string `json:"date"`
}
