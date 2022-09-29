package main

import (
	"cadena/blocke/blocke"
	"cadena/blocke/handler"
	"cadena/blocke/handler/routes"
	"log"
	"net/http"
)

// func Router(api *handler.Api) http.Handler {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/", api.HandleNewBlock)

// 	return mux
// }

func main() {

	api := handler.Api{
		Blocke:  blocke.NewBlocke(),
		Blockes: blocke.NewBlockes(),
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: routes.Router(&api),
	}

	log.Println("Server start at http://localhost:8081/")

	log.Fatal(srv.ListenAndServe())
}
