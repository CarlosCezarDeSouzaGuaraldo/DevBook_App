package controllers

import "net/http"

//CreateUser is the CREATE method
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))
}

//GetUsers is the method to GET all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

//GetUser is the method to GET an unique user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get an unique user"))
}

//UpdateUser is a method to UPDATE an unique user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an unique user"))
}

//DeleteUser is a method to DELETE an unique user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Remove an unique User"))
}
