package models

import (
	"fmt"
)

type Response struct {
	Code        string 	   `json:"code"`   	 	
	Message		string 	   `json:"message"`
}

func (w *Response) Error() string {
    return fmt.Sprintf(w.Code,w.Message)
}

func Wrap(code string, message string) *Response {
    return &Response{
        Code		: code,
		Message		: message,
    }
}