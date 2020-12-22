package interfaces

import(
	"../models"
)

type UserRepository interface{
	InsertNewUser(User models.Users) (bool, error)
	GetUserbyUsername(Username string) (users []models.Users, err error)
}