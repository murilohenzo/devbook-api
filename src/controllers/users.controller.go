package controllers

import (
	"api_resources/src/models"
	"api_resources/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Create create new user
func Create(w http.ResponseWriter, r *http.Request)  {
	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, "Error ao cadastra novo usuario", http.StatusUnprocessableEntity)
		return
	}
	var user models.User

	if erro = json.Unmarshal(request, &user); erro != nil {
		http.Error(w, "Error ao cadastra novo usuario", http.StatusBadRequest)
		return
	}

	err := repositories.Create(&user)
	if err != nil {
		return 
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return 
	}
}

// FindAll find all users
func FindAll(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Find all users"))
}

// FindById find user by id
func FindById(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Find user by id"))
}

// Update update user by id
func Update(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Updated user by id"))
}

// Delete delete user by id
func Delete(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Delete user by id"))
}