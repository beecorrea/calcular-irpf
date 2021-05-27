package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (app *App) Routes(){
	router := mux.NewRouter();
	router.HandleFunc("/salary", CalcSalaryHandler);
	
	app.Router = router;
}

func (app *App) Run(){
	// Set cors handler.
	handler := cors.Default().Handler(app.Router);

	log.Print("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}