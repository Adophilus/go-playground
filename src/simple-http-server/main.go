package main

import (
  "net/http"
)

func index (res http.ResponseWriter, req *http.Request) {
  res.WriteHeader(http.StatusTeapot)
}

func main () {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}
