package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloMuxHandler(
	w http.ResponseWriter,
	r *http.Request) {
	w.Write([]byte("Hello gorilla/mux!\n"))
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
	}

	return
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	r.HandleFunc("/", helloMuxHandler)

	/* matchers */
	r.HandleFunc("/products", ProductHandler).
		Host("example.com").   /* Only match a specific domain */
		Methods("GET", "PUT"). /* Only match GET and PUT methods */
		Schemes("http")        /* Only match the http scheme */

	log.Fatal(http.ListenAndServe(":8080", r))

}
