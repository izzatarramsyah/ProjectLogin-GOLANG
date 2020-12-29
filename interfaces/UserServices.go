package interfaces

import(
	"github.com/izzatarramsyah/ProjectLogin-GOLANG/models"
)

type UserServices interface{
	UserRegistration(User models.Users)  (bool, error)
	GetUser(Username string) (users []models.Users, err error)
	GetCheckPassword(Username string, Password string) (bool, error)
}