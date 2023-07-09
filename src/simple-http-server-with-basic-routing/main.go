package main

import (
  "net/http"
)

func index (res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Welcome to the index page!"))
}

func login (res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Welcome to the login page!"))
}

func main () {
  http.HandleFunc("/", index)
  http.HandleFunc("/login", login)
  http.ListenAndServe(":8080", nil)
}
