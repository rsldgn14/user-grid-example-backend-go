package controllers

import (
	"deneme/pgk/models"
	"deneme/pgk/utils"
	"encoding/json"

	"fmt"

	"github.com/gorilla/mux"

	"net/http"

	"strconv"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	NewUsers := models.GetAllUser()
	res, _ := json.Marshal(NewUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")

	}

	userDetails, _ := models.GetUserById(ID)

	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}

	utils.ParseBody(r, CreateUser)

	u := CreateUser.CreateUser()

	res, _ := json.Marshal(u)

	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")

	}

	user := models.DeleteUser(ID)

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var updateUser = &models.User{}

	utils.ParseBody(r, updateUser)

	vars := mux.Vars(r)

	userId := vars["userId"]

	ID, err := strconv.ParseInt(userId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")

	}

	userDetails, db := models.GetUserById(ID)
	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	db.Save(&userDetails)

	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}
