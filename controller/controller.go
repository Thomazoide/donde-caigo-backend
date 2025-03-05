package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/Thomazoide/donde-caigo-backend/docs"
	"github.com/Thomazoide/donde-caigo-backend/middleware"
	"github.com/Thomazoide/donde-caigo-backend/structs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(h apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			fmt.Println(err)
			response := &structs.ApiResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "BAD REQUEST",
				Error:      err.Error(),
			}
			WriteJSON(w, http.StatusBadRequest, response)
		}
	}
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) RunServer() {
	router := mux.NewRouter()
	middleware.EnableCORS(router)
	router.Use(middleware.MiddlewareCORS)
	router.Use(middleware.MiddleWareCookieConsumer)
	router.HandleFunc("/auth", makeHTTPHandlerFunc(s.handleAuth))
	router.HandleFunc("/cuenta", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/cuenta/{id}", makeHTTPHandlerFunc(s.handleAccountWithParams))
	router.HandleFunc("/publicaciones", makeHTTPHandlerFunc(s.handlePost))
	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(s.listenAddr, router)
}
