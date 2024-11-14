package models

type Person struct {
	Id    int64  `json:"id" fake:"skip"`
	Name  string `json:"name" fake:"{firstname}"`
	Email string `json:"email" fake:"{email}"`
	Phone string `json:"phone" fake:"{phone}"`
}
