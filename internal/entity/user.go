package entity

import "time"

type User struct {
	Id          uint64    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Sex         string    `json:"sex"`
	City        string    `json:"city"`
	DateBirth   string    `json:"dateBirth"`
	DateCreated time.Time `json:"dateCreated"`
}
