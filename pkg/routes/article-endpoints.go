package routes

import (
	controller "GoWebServiceWithCRUD/pkg/controllers"
	//"GoWebServiceWithCRUD/pkg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ArticleEndPointsRequestHandler() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc

	myRouter.HandleFunc("/", controller.HomePage)
	myRouter.HandleFunc("/all", controller.ReturnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", controller.ReturnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article", controller.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", controller.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/update/{id}", controller.UpdateArticle).Methods("PUT")

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
