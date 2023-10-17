package main

import (
	"fmt"
	"go-auth/handler"
	"go-auth/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load env from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loaded .env file")
	}

	// initialize db connection
	db, err := repository.NewDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// initialze repository
	userRepository := repository.NewUserRepository(db)

	// initialze route

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success":"true"}`))
	})

	r.HandleFunc("login", func(w http.ResponseWriter, r *http.Request) {
		handler.AuthHandler(w, r, userRepository)
	}).Methods("POST")

	r.HandleFunc("register", func(w http.ResponseWriter, r *http.Request) {
		handler.RegisterHandler(w, r, userRepository)
	}).Methods("POST")
	http.Handle("/", r)
	// listen and serve
	PORT := os.Getenv("PORT")

	fmt.Print("server running on port : ", PORT)
	http.ListenAndServe(":"+PORT, nil)

}
