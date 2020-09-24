package routes

import (
	controller "GoWebServiceWithCRUD/pkg/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CustomerEndPointsRequestHandler() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc

	myRouter.HandleFunc("/customer/getall", controller.GetAllCustomers).Methods("GET")
	myRouter.HandleFunc("/customer/{name}", controller.GetSpecificCustomer).Methods("GET")
	myRouter.HandleFunc("/customer/createnew", controller.CreateNewCustomer).Methods("POST")
	myRouter.HandleFunc("/customer/delete/{name}", controller.DeleteCustomer).Methods("DELETE")
	myRouter.HandleFunc("/customer/update/{name}", controller.UpdateCustomer).Methods("PUT")



	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
