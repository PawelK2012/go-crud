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
	router.HandleFunc("/notes", makeHTTPHandleFunc(s.handleMenu))
	log.Println("server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleMenu(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		log.Println("GET note API")
		return s.handleGetMenuByID(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateNote(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleGetMenuByID(w http.ResponseWriter, r *http.Request) error {

	menuId := 1
	menu, err := s.store.GetNoteById(menuId)
	if err != nil {
		formattedErr := fmt.Errorf("menu not found %s", err)
		log.Print(formattedErr)
		return formattedErr

	}
	return WriteJSON(w, http.StatusOK, menu)
}

func (s *APIServer) handleCreateNote(w http.ResponseWriter, r *http.Request) error {
	// note := models.Note{}
	note := new(models.Note)
	if err := json.NewDecoder(r.Body).Decode(note); err != nil {
		return err
	}

	newNote := models.NewNote(note.Author, note.Title, note.Desc, note.Tags)
	if err := s.store.CreateNote(newNote); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, note)
}
