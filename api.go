package main // main package

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WriteJSON function is for putting v into JSON
func WriteJSON(w http.ResponseWriter, status int, v any) error {
    w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// APIFunc is the type for an api function
type APIFunc func(http.ResponseWriter, *http.Request) error

// APIError is a struct and type that holds the error message
type APIError struct {
	Error string
}

// makeHTTPHandleFunc wraps API functions so that the mux library can handle them since router.HandleFunc needs a parameter of type http.HandlerFunc
func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

// APIServer is a struct in charge of holding some server data
type APIServer struct {
    store Storage
	listenAddr string
}

// NewAPIServer function is in charge of initializing our server
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
        store Storage,
	}
}

// Run funtion is in charge of starting the server
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

    router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// handleAccount funtion is in charge of other handling the requests initally depending on the request method
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error{
	switch method := r.Method; method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("Not allowed method %s", method)
	}
}

// handleGetAccount function is in charge of getting user accounts
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)

    //db.get(id)
    return WriteJSON(w, http.StatusOK, vars)
}

// handleCreateAccount function is in charge of creating user accounts
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleDeleteAccount functino is in charge of deleting user accounts
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleProgress function is in charge of handling the progress the User makes in the tutorial
func (s *APIServer) handleProgress(w http.ResponseWriter, r *http.Request) error {
	return nil
}
