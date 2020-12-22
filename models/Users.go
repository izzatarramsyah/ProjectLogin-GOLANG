package models

import (
	"time"
	"fmt"
)	

type Users struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Hash       	   string `json:"hash"`
	Email          string `json:"email"`
	Created_date   time.Time `json:"created_date"`
	Created_by     string `json:"created_by"`
	Lastupd_date   time.Time `json:"lastupd_date"`
	Lastupd_by 	   string `json:"lastupd_by"`
	Status 		   string `json:"status"`
	Url 		   string `json:"url"`
}

type ResponseUsers struct {
	Code        string 	   `json:"code"`   	 	
	Message		string 	   `json:"message"`
	Object		[]Users    `json:"object"`
}

func (w *ResponseUsers) ErrorUsers() string {
    return fmt.Sprintf(w.Code,w.Message, w.Object)
}

func WrapUsers(code string, message string, object []Users) *ResponseUsers {
    return &ResponseUsers{
        Code		: code,
		Message		: message,
		Object		: object,
    }
}