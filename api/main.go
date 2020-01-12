package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"api/db"
	"github.com/gorilla/mux"

	"api/controller"
)

func main() {

	log.SetOutput(os.Stdout)
	ctx := context.Background()

	mongo, err := db.ConfigureDB(ctx)
	if nil != err {
		panic("Error connecting to mongo")
	}

	router := mux.NewRouter()

	repository := db.NewEntryRepository(mongo)

	if nil != err {
		panic("Error building repository")
	}

	controller := controller.NewEntryController(repository) 

	router.HandleFunc("/entries", controller.GetEntriesHandler).Methods("GET")

	router.HandleFunc("/entries", controller.PostEntriesHandler).Methods("POST")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "This is a website served by a Go HTTP server.")
	// })

	// http.HandleFunc("/oldentry", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("entries!")
	// 	err = repository.Create()
	// 	// err = repository.GetAll()
	// 	fmt.Fprintf(w, "This is the entry route!")
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.Handle("/", )

	// test := "my test"

	fmt.Println("Server started successfully!")
	http.ListenAndServe(":3001", router)
}

