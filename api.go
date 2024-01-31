package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PawelK2012/go-crud/models"
	"github.com/PawelK2012/go-crud/repository"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	repo       repository.Repository
}

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, repo repository.Repository) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		repo:       repo,
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
	router.HandleFunc("/note", makeHTTPHandleFunc(s.handleNote))
	router.HandleFunc("/notes", makeHTTPHandleFunc(s.handleNotes))
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		panic(err)
	}
	log.Println("server running on port: ", s.listenAddr)
}

func (s *APIServer) handleNotes(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		log.Println("GET notes API")
		return s.handleGetAllNotes(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleNote(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		log.Println("GET note API")
		return s.handleGetNoteByID(w, r)
	}

	if r.Method == "POST" {
		log.Println("POST note API")
		return s.handleCreateNote(w, r)
	}
	//using PUT for simplicity
	if r.Method == "PUT" {
		log.Println("PUT note API")
		return s.handleUpdateNote(w, r)
	}

	if r.Method == "DELETE" {
		log.Println("DELETE note API")
		return s.handleDeleteNote(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleGetNoteByID(w http.ResponseWriter, r *http.Request) error {

	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed parsing ID", err)
		panic(err)
	}
	menu, err := s.repo.GetById(r.Context(), i)
	if err != nil {
		log.Println("get note by ID failed with error:", err)
		return err

	}
	return WriteJSON(w, http.StatusOK, menu)
}

// add r.Body validation
func (s *APIServer) handleCreateNote(w http.ResponseWriter, r *http.Request) error {
	note := &models.Note{}
	if err := json.NewDecoder(r.Body).Decode(note); err != nil {
		log.Println("failed decoding user payload", err)
		return err
	}

	newNote := models.NewNote(note.Author, note.Title, note.Desc, note.Tags)
	n, err := s.repo.Create(r.Context(), *newNote)
	if err != nil {
		log.Println("creating note failed with error:", err)
		return err
	}
	newNote.Id = n

	return WriteJSON(w, http.StatusOK, newNote)
}

func (s *APIServer) handleUpdateNote(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	note := &models.Note{}
	if err := json.NewDecoder(r.Body).Decode(note); err != nil {
		log.Println("failed decoding user payload", err)
		return err
	}

	n, err := s.repo.Update(r.Context(), id, *note)
	if err != nil {
		log.Println("updating note failed with error:", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, n)
}

func (s *APIServer) handleDeleteNote(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed parsing ID", err)
		panic(err)
	}

	delID, err := s.repo.Delete(r.Context(), i)
	if err != nil {
		log.Println("deleting note failed with error:", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, delID)
}

func (s *APIServer) handleGetAllNotes(w http.ResponseWriter, r *http.Request) error {
	n, err := s.repo.GetAll(r.Context())
	if err != nil {
		log.Println("get all notes failed with error:", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, n)
}
