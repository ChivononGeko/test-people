package transport

import (
	"net/http"
	_ "test-people/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(handler *PersonHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /person", handler.CreatePerson)
	mux.HandleFunc("DELETE /person", handler.DeletePerson)
	mux.HandleFunc("PUT /person", handler.UpdatePerson)
	mux.HandleFunc("GET /person", handler.GetPersonByFilters)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.Handle("/swagger/index.html", httpSwagger.WrapHandler)

	return mux
}
