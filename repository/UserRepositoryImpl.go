package repository

import (
	"log"
	"../models"
	_util "../util"
	"database/sql"
	"time"
	"context"
)

var err error

type UserRepository struct{
	Conn *sql.DB
	Ctx context.Context
}

func (repository *UserRepository) InsertNewUser(Users models.Users) (bool, error) {
	log.Println("Endpoint Hit: InsertNewUser")
	hash,_ := _util.HashPassword(Users.Password)

	ctx, cancel := context.WithTimeout(repository.Ctx, time.Second*2)
	defer cancel()

	createDt := time.Now()

	stmt, err := repository.Conn.PrepareContext(ctx,"INSERT INTO USERS (USERNAME, PASSWORD, HASH, EMAIL, CREATED_DATE, CREATED_BY, STATUS, URL) VALUE (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("error PrepareContext : " ,err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, Users.Username, Users.Password, hash, Users.Email, createDt, Users.Created_by, Users.Status, Users.Url)

	if err != nil {
		log.Println("error ExecContext : " ,err)
		return false,err
	} else {
		rows, _ := result.RowsAffected()
		if rows != 0 {
			return true,nil
		}else{
			return false,nil
		}
	}
}

func (repository *UserRepository) GetUserbyUsername(Username string) (users []models.Users, err error){
	log.Println("Endpoint Hit: GetUserByUsername : " + Username)

	ctx, cancel := context.WithTimeout(repository.Ctx, time.Second*2)
	defer cancel()
	
	row,err := repository.Conn.QueryContext(ctx, "SELECT USERNAME, PASSWORD, HASH, EMAIL, CREATED_DATE, CREATED_BY, STATUS, URL FROM USERS WHERE USERNAME = ?",Username)
	defer row.Close()

	for row.Next(){
		var user models.Users
		if err := row.Scan(&user.Username, &user.Password, &user.Hash, &user.Email, &user.Created_date, &user.Created_by, &user.Status, &user.Url); err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil 
}