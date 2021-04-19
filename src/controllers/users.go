package controllers

import (
	"net/http"
)

//CreateUser insert one user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criating user!"))
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
