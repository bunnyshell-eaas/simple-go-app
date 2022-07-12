package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Author *Author
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Author: &Author{Firstname: "Johnny", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Author: &Author{Firstname: "Michael", Lastname: "Witty"}})
	books = append(books, Book{ID: "3", Author: &Author{Firstname: "Ion", Lastname: "Creanga"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/", indexHtml).Methods("GET")

	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Print(w, "This is my website!\n")

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}
