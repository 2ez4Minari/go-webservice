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

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Customers Endpoint Hit")

	db, err := gorm.Open(config.DbType, config.SqlConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()

	var customers []models.Customer
	db.Find(&customers)
	fmt.Println("{}", customers)

	json.NewEncoder(w).Encode(customers)
}

func GetSpecificCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Specific Customer Endpoint Hit")

	db, err := gorm.Open(config.DbType, config.SqlConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var customer models.Customer
	db.Where("customer_name = ?", name).Find(&customer)
	db.Find(&customer)

	fmt.Println("{}", customer)
	json.NewEncoder(w).Encode(customer)
}

func CreateNewCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create New Customer Endpoint Hit")

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
	var customer models.Customer
	json.Unmarshal(reqBody, &customer)

	db.Create(&models.Customer{CustomerName: customer.CustomerName, CustomerAge: customer.CustomerAge, CustomerAddress: customer.CustomerAddress,
		CustomerEmail: customer.CustomerEmail, IsActive: customer.IsActive})
	fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Customer Endpoint Hit")

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

	var customer models.Customer
	db.Where("customer_name = ?", name).Find(&customer)
	db.Delete(&customer)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Customer Endpoint Hit")

	db, err := gorm.Open(config.DbType, config.SqlConnectionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedCustomer models.Customer
	json.Unmarshal(reqBody, &updatedCustomer)

	var customerRecord models.Customer
	db.Where("customer_name = ?", name).Find(&customerRecord)

	customerRecord.CustomerName = updatedCustomer.CustomerName
	customerRecord.CustomerAge = updatedCustomer.CustomerAge
	customerRecord.CustomerAddress = updatedCustomer.CustomerAddress
	customerRecord.CustomerEmail = updatedCustomer.CustomerEmail
	customerRecord.IsActive = updatedCustomer.IsActive

	db.Save(&customerRecord)
	fmt.Fprintf(w, "Successfully Updated User")
}
