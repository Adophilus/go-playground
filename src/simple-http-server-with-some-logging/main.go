package main

import (
	"log"
	"net/http"
)

func index (res http.ResponseWriter, req *http.Request) {
  log.Print("A request was made to /")
  res.WriteHeader(http.StatusTeapot)
}

func main () {
  http.HandleFunc("/", index)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatalf("Failed to start up server. Reason = %v", err)
    return
  }
}
