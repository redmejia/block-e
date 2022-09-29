package routes

import (
	"cadena/blocke/handler"
	"net/http"
)

func Router(a *handler.Api) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/new/block", a.HandleNewBlock)
	mux.HandleFunc("/cadena", a.BlockesHandler)
	return mux
}
