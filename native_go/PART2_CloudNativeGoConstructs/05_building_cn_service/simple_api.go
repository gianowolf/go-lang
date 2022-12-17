package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var myMap = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrorNoSuchKey = errors.New("no such key")

/* * * * * * * * * * * * * * * * */
//    API RESTFUL METHODS        //
/* * * * * * * * * * * * * * * * */

func Get(key string) (string, error) {
	value, ok := capitals[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}
func Put(key, value string) error {
	capitals[key] = value
	return nil
}

func Delete(key string) error {
	delete(capitals, key)
	return nil
}

/* * * * * * * * * * * * * * * * */
//           HANDLERS            //
/* * * * * * * * * * * * * * * * */

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	Delete(key)

	w.WriteHeader(http.StatusCreated)
}

/* * * * * * * * * * * * * * * * */
//           MAIN                //
/* * * * * * * * * * * * * * * * */

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/{key}", keyValueGetHandler).
		Methods("GET")
	r.HandleFunc("/v1/{key}", keyValuePutHandler).
		Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).
		Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
