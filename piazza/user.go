package piazza

import (

)


type UserType int

const (
	InvalidUser UserType = iota
	NormalUser
	AdminUser
)

type User struct {
	Id   PiazzaId
	UserType UserType
	Name string
}
