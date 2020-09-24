package routes

import (
	controller "GoWebServiceWithCRUD/pkg/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UserEndPointsRequestHandler() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", controller.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", controller.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", controller.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user", controller.NewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
