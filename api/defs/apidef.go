package defs


// requests


type UserCredential struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}


// Data model

type User struct {
	Id int
	LoginName string
	Pwd string
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comments struct {
	Id string
	VideoInfo string
	AuthorId int
	Comment string
}