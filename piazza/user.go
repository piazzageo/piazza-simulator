package piazza

import (

)

type UserId int64

var nextUserId MessageId = 1

type UserType int

const (
	InvalidUser UserType = iota
	NormalUser
	AdminUser
)

type User struct {
	Id   UserId
	UserType UserType
	Name string
}
