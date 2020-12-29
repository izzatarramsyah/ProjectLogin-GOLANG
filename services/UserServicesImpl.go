package services

import (
	"log"
	"github.com/izzatarramsyah/ProjectLogin-GOLANG/models"
	"github.com/izzatarramsyah/ProjectLogin-GOLANG/interfaces"
	_util "github.com/izzatarramsyah/ProjectLogin-GOLANG/util"
)

var err error

type UserServices struct {
	Repo interfaces.UserRepository
}

func (service * UserServices) UserRegistration(User models.Users) (bool, error){
	log.Println("Endpoint Hit: UserRegistration")

	res,err := service.Repo.InsertNewUser(User)
	return res,err
}

func (service * UserServices) GetUser(Username string) (users []models.Users, err error){
	log.Println("Endpoint Hit: GetUser")

	res,err := service.Repo.GetUserbyUsername(Username)
	return res,err
}

func (service * UserServices) GetCheckPassword(Username string, Password string) (bool, error){
	log.Println("Endpoint Hit: GetCheckPassword")

	res,_ := service.Repo.GetUserbyUsername(Username)
	if res!= nil {
		hash := res[0].Hash
		match := _util.CheckPasswordHash(Password, hash)

		if match == true{
			return true,nil
		}else{
			return false,nil
		}

	}else{
		return false,nil
	}

	
}