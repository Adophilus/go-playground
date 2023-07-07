package main

import (
	"crypto/subtle"
	"log"
	"net/http"
)

const (
  ADMIN_USERNAME = "admin"
  ADMIN_PASSWORD = "password"
)

func index (res http.ResponseWriter, req *http.Request) {
  res.WriteHeader(http.StatusTeapot)
}

func basicAuth (handler http.HandlerFunc, realm string) http.HandlerFunc {
  return func (res http.ResponseWriter, req *http.Request) {
    username, password, ok := req.BasicAuth()
    if !ok || subtle.ConstantTimeCompare([]byte(username), []byte(ADMIN_USERNAME)) != 1 ||subtle.ConstantTimeCompare([]byte(password), []byte(ADMIN_PASSWORD)) != 1 {
      res.Header().Set("WWW-Authenticate", `realm="` + realm + `"`)
      res.WriteHeader(http.StatusForbidden)
      res.Write([]byte("You are not authorized to access this page!"))
    } else {
      handler(res, req)
    }
  }
}

func protected (res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Welcome!"))
}

func main () {
  http.HandleFunc("/", index)
  http.HandleFunc("/auth", basicAuth(protected, "Please authenticate"))
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatalf("Failed to start server. Reason: %v", err)
  } else {
    log.Print("Running server on port 8080 🚀")
  }
}
