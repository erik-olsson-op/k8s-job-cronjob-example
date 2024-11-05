package models

type Person struct {
	Id    int64  `fake:"skip"`
	Name  string `fake:"{firstname}"`
	Email string `fake:"{email}"`
	Phone string `fake:"{phone}"`
}
