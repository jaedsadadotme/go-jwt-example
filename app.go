package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App export
type App struct {
	Router *mux.Router
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()

	// auth
	authrouter := app.Router.
		PathPrefix("/auth/").
		Subrouter()
	authrouter.HandleFunc("/login", loginHandler).Methods("POST")

	// app
	appRouter := app.Router.
		PathPrefix("/api/v1/").
		Subrouter()
	appRouter.Use(Authenticate)
	appRouter.HandleFunc("/", helloWorldHandler).Methods("GET")
}

func (app *App) run() {
	port := ":8080"
	fmt.Println("===== start =====")
	fmt.Println("http://localhost:" + port)
	fmt.Println("-----------------")
	log.Fatal(http.ListenAndServe(port, app.Router))
}
