package main

import (
	// "io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func index (res http.ResponseWriter, req *http.Request) {
  // io.WriteString(res, "Hello there 👋")
  res.Write([]byte("Hello there 👋"))
}

func main () {
  mux := http.NewServeMux()
  mux.HandleFunc("/", index)
  err := http.ListenAndServe(":8080", handlers.CompressHandler(mux))
  log.Print("Running server...")
  if err != nil {
    log.Fatalf("Failed to start server. Reason: %v", err)
    return
  }
}
