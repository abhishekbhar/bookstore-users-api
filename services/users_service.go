package services


import (
	"github.com/abhishekbhar/bookstore-users-api/domain/users"
)



func CreateUser(user users.User) (*users.User, error){
	return &user,nil
}