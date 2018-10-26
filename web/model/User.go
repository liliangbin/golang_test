package model

type User struct {
	ID   int `gorm:"primary_key"`
	Name string
	Info string
}

type Message struct {
	ID           int `grom:"primary_key"`
	UserId       int
	Message      string
	Introduction string
}
