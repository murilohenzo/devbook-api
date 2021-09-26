package controllers

import (
	"api_resources/src/json_error"
	"api_resources/src/models"
	"api_resources/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HandleException struct {
	Error string `json:"error"`
}

func ErrorRequestHandle(text string) (errorRequestHandle HandleException) {
	errorRequestHandle = HandleException{
		Error: text}
	return
}

// Create create new user
func Create(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		json_error.JSONError(w, ErrorRequestHandle(err.Error()), http.StatusBadRequest)
		return
	}
	var user models.User

	if err = json.Unmarshal(request, &user); err != nil {
		json_error.JSONError(w, ErrorRequestHandle(err.Error()), http.StatusBadRequest)
		return
	}

	err = repositories.Create(&user)
	if err != nil {
		json_error.JSONError(w, ErrorRequestHandle(err.Error()), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		json_error.JSONError(w, ErrorRequestHandle(err.Error()), http.StatusInternalServerError)
		return
	}
}

// FindAll find all users
func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all users"))
}

// FindById find user by id
func FindById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user by id"))
}

// Update update user by id
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updated user by id"))
}

// Delete delete user by id
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user by id"))
}
