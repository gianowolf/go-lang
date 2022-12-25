package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

type EventType byte

type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key, value string)
}

const (
	_                     = iota
	EventDelete EventType = iota
	EventPut
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrorNoSuchKey = errors.New("no such key")

/* TRANSACTION LOGGER INTERFACE */
/*	the transaction log file mantains a history of mutating changes
/*	Makes it possible to replay the transactions to reconstruct the service's state
/*	Ordered list of mutating events */
/* * * * * * * * * * * * * * * * * * * */
type FileTransactionLogger struct {
	events      chan<- Event /* W-Only CH for seding EV */
	errors      <-chan error /* R-Only CH for receive ERR */
	lastSeuence uint64       /* last used event sequence number */
	file        *os.File     /* location of the transaction log */
}

func (l *FileTransactionLogger) WritePut(key, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *FileTransactionLogger) Err() <-chan error {
	return l.errors
}

func NewFileTransactionLogger(filename string) (TransactionLogger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file : %w", err)
	}
	return &FileTransactionLogger{file: file}, nil
}

/* EVENTS */
/* Internal representation of events to asynchronous operations */
/* * * * * */

type Event struct {
	Sequence  uint64    // Sequence Number
	EventType EventType // A constant that represents the different type of event
	Key       string    // Vey Affected
	Value     string    // Value of PUT transactions
}

/* * * * * * * * * * * */
/*  REST API METHODS   */
/* * * * * * * * * * * */
func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}
func Put(key, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil
}

func Delete(key string) error {
	delete(store.m, key)
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
