package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/izzatarramsyah/ProjectLogin-GOLANG/controller"
	"github.com/izzatarramsyah/ProjectLogin-GOLANG/repository"
	"github.com/izzatarramsyah/ProjectLogin-GOLANG/services"
	_configuration "github.com/izzatarramsyah/ProjectLogin-GOLANG/config"
	"database/sql"
	"github.com/go-chi/chi"
	"context"
	_ "github.com/go-sql-driver/mysql" 
	"os"
)

func main() {

	config := _configuration.NewConfiguration()
	config.LoadConfigurationFromFile(getFilePathConfigEnvirontment())
	dbHost := config.GetValue(`database.host`)
	dbPort := config.GetValue(`database.port`)
	dbUser := config.GetValue(`database.user`)
	dbPass := config.GetValue(`database.pass`)
	dbName := config.GetValue(`database.name`)
	connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True")

	log.Println(connection)
	sqlConn,err := sql.Open("mysql", "root:password@tcp(godockerDB)/dev_apps")
	c := context.Background()

	if(err != nil){
		log.Println("Connect Failed")
	}else{
		log.Println("Connect Success")
	}

	userRepository := &repository.UserRepository{Conn : sqlConn, Ctx : c}
	userService := &services.UserServices{Repo : userRepository}
	userController := controller.UserController{Service : userService}


	r := chi.NewRouter()
	r.Post("/User/Registration", userController.UserRegis)
	r.Post("/User/ListUsers", userController.GetUsers)
	r.Post("/User/UserLogin", userController.UserLogin)
	r.Get("/User/Test", userController.Test)

	http.ListenAndServe(":8080", r)
}

func getFilePathConfigEnvirontment() string {
	env := "dev"
	if len(os.Args) >= 2 && os.Args[1] != "" {
		log.Println(os.Args[1])
		env = os.Args[1]
	} else {
		log.Println("you must define environtment : 'go run main.go dev'")
	}
	switch env {
	case "dev":
		return "config-dev.json"
	case "staging":
		return "config-staging.json"
	case "prod":
		return "config-prod.json"
	}
	return "config-dev.json"
}