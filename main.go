package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
//Defind data model
type Book struck{
	ID		string		`json:"id"`
}

func main(){
//Init Router 
router := mux.NewRouter()

//Handler Function/End Points
router.HandleFunc("/api/books",getBooks).Methods("GET")
router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
router.HandleFunc("/api/books",createBooks).Methods("POST")
router.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
router.HandleFunc("/api/books/{id}",deleteBooks).Methods("DELETE")

//Listen on port 
log.Fatal(http.ListenAndServe(":8000",router))



}