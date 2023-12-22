package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PawelK2012/go-crud/models"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type ApiError struct {
	Error string `json:"error"`
}

// type Menu struct {
// 	Breakfast string `json:"breakfast"`
// 	Lunch     string `json:"lunch"`
// 	Dinner    string `json:"dinner"`
// }

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
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
		log.Println("orders API")
		return s.handleGetOrderByID(w, r)
	}

	// if r.Method == "POST" {
	// 	return s.handleCreateAccount(w, r)
	// }
	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleGetOrderByID(w http.ResponseWriter, r *http.Request) error {
	menu := models.Menu{Breakfast: "Scrambled eggs", Lunch: "Burgers && Fires", Dinner: "Coco jubmbo"}
	return WriteJSON(w, http.StatusOK, menu)
}
