package controllers

import (
	"GoWebServiceWithCRUD/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	models.Articles = []models.Article{
		models.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		models.Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(models.Articles)
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	models.Articles = append(models.Articles, article)

	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range models.Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			models.Articles = append(models.Articles[:index], models.Articles[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(models.Articles)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {

	// get the body of our PUT request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle models.Article
	json.Unmarshal(reqBody, &newArticle)

	// update our global Articles array to include
	// our updated Article
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	// we will need to extract the `id` of the article we
	// wish to update
	id := vars["id"]

	// we then need to loop through all our articles
	for i := range models.Articles {
		// if our id path parameter matches one of our
		// articles
		if models.Articles[i].Id == id {
			// updates our Articles array to remove the
			// article
			models.Articles[i].Title = newArticle.Title
			models.Articles[i].Content = newArticle.Content
			models.Articles[i].Desc = newArticle.Desc
		}
	}
	json.NewEncoder(w).Encode(models.Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range models.Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
