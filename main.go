package main

import (
	"log"
	"net/http"
	"os"

	"bodra.com.br/go-recipe-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Auth routes
	router.HandleFunc("/signin", controllers.Signin)
	router.HandleFunc("/refreshtoken", controllers.Signin)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		log.Println("$PORT not set. Falling back to default " + port)
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
