package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Thomazoide/donde-caigo-backend/middleware"
	"github.com/Thomazoide/donde-caigo-backend/structs"
	"github.com/gorilla/mux"
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
		if err := h(w, r); err != nil {
			response := &structs.ApiResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "BAD REQUEST",
				Error:      err,
			}
			WriteJSON(w, http.StatusBadRequest, response)
			return
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
	router.HandleFunc("/auth", makeHTTPHandlerFunc(s.handleAuth))
	router.HandleFunc("/cuenta", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/cuenta/{id}", makeHTTPHandlerFunc(s.handleAccountWithParams))
	router.HandleFunc("/publicaciones", makeHTTPHandlerFunc(s.handlePost))
	http.ListenAndServe(s.listenAddr, router)
}
