package controllers

import (
	"GoWebServiceWithCRUD/pkg/config"
	"GoWebServiceWithCRUD/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(config.DbType, config.SqlConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()

	var users []models.User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open(config.DbType, config.SqlConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Println("Connected to database")
	}
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	reqBody, _ := ioutil.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(reqBody, &user)

	db.Create(&models.User{Name: user.Name, Email: user.Email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(config.DbType, config.SqlConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Println("Connected to database")
	}
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user models.User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(config.DbType, config.SqlConnectionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user models.User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}
