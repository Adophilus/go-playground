package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getRequestHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "You've hit the GET request handler")
}

func postRequestHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "You've hit the POST request handler")
}

func pathVariableRequestHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	fmt.Fprintf(res, "Welcome %s", name)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getRequestHandler).Methods("GET")
	router.HandleFunc("/", postRequestHandler).Methods("POST")
	router.Handle("/{name}", http.HandlerFunc(pathVariableRequestHandler)).Methods("GET", "POST", "PUT")
	http.ListenAndServe(":8080", router)
}
