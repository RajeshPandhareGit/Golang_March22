package main

import (
	"encoding/json"
	"log"
	"net/http"
	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct{
	ID 		string 		`json:"id"`
	Isbn 	string 		`json:"isbn"`
	Title	string 		`json:"title"`
	Author	*Author		`json:"author"`
}

//Author struct
type Author struct{
	Firstname 	string 	`json:"firstname"`
	LastName	string 	`json:"lastname"`
}

var books []Book

//init book variable as slice 

//Implementation of getBooks function
func getBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter,r *http.Request){

}
func deleteBooks(w http.ResponseWriter,r *http.Request){

}
func updateBooks(w http.ResponseWriter,r *http.Request){

}
func insertBooks(w http.ResponseWriter,r *http.Request){

}

func main(){
	//Initialization of Router
	r:= mux.NewRouter()

	books = append (books, Book{ID:"1", Isbn:"448743", Title:"Book1", Author: &Author{
		Firstname:"John",LastName:"Doe",}})
	books = append (books, Book{ID:"2", Isbn:"448744", Title:"Book2", Author: &Author{
			Firstname:"Mike",LastName:"Billord",}})

	//Define Route handler . End Points for API.
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books", insertBooks).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))



}