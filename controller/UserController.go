package controller

import (
	"log"
	"../models"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../interfaces"
)

var err error

type UserController struct{
	Service interfaces.UserServices
}

func (controller *UserController) UserRegis(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: UserRegis")
	commonResponse := models.Wrap("200", "Success")

	var user models.Users
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(w, "Kindly enter data with the event Consta and Value only in order to update")
	}

	jsonString := string(reqBody)
	json.Unmarshal([]byte(jsonString), &user)

	_,errr := controller.Service.UserRegistration(user)

	if errr != nil {
		message,_ := "Failed. Error Registration ", errr
		commonResponse = models.Wrap("500", message)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(commonResponse)
}

func (controller *UserController) GetUsers(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: GetUsers")

	var user models.Users
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(w, "Kindly enter data with the event Consta and Value only in order to update")
	}

	jsonString := string(reqBody)
	json.Unmarshal([]byte(jsonString), &user)
	log.Println("Username : ",user.Username)
	
	users,err := controller.Service.GetUser(user.Username)

	commonResponse := models.WrapUsers("200", "Success", users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commonResponse)
}

func (controller *UserController) UserLogin(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: UserLogin")

	var user models.Users
	var Message string

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(w, "Kindly enter data with the event Consta and Value only in order to update")
	}

	jsonString := string(reqBody)
	json.Unmarshal([]byte(jsonString), &user)
	log.Println("Username : ",user.Username)
	log.Println("Password : ",user.Password)

	res,err := controller.Service.GetCheckPassword(user.Username, user.Password)

	if res == true{
		Message = "User Is Valid "
	}else{
		Message = "User Is Not Valid "
	}

	commonResponse := models.Wrap("200", Message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commonResponse)
}