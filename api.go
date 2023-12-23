package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Store
}

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, store Store) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", makeHTTPHandleFunc(s.handleMenu))
	log.Println("server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleMenu(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		log.Println("menu API")
		return s.handleGetMenuByID(w, r)
	}

	// if r.Method == "POST" {
	// 	return s.handleCreateAccount(w, r)
	// }
	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleGetMenuByID(w http.ResponseWriter, r *http.Request) error {
	// for simplicity of this example we have only 1 menu
	menuId := 1
	menu, err := s.store.GetMenu(menuId)
	if err != nil {
		formattedErr := fmt.Errorf("menu not found %s", err)
		log.Print(formattedErr)
		return formattedErr

	}
	return WriteJSON(w, http.StatusOK, menu)
}
