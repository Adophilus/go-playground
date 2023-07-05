package main

import (
  "crypto"
  "net/http"
)

func index (res http.ResponseWriter, req *http.Request) {
  res.WriteHeader(http.StatusTeapot)
}

func auth (res http.ResponseWriter, req *http.Request) {

}

func main () {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}
