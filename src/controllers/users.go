package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CreateUser insert one user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Connect error:", err)
	}

	repository := repositories.NewUserRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", userId)))
}

//FindAllUsers find all users in database
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding all users!"))
}

//FindUser find one user in database
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user!"))
}

//UpdateUser update user infos in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

//DeleteUser delete user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
