package main

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	capitals       = make(map[string]string)
	ErrorNoSuchKey = errors.New("no such key")
)

func Put(key, value string) error {
	capitals[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, ok := capitals[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) error {
	delete(capitals, key)
	return nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo!\n"))
}

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1", defaultHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
