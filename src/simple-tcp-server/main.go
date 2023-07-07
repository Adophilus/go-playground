package main

import (
  "log"
  "net"
)

func main () {
  listener, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatal("Failed to start server. Reason:", err)
    return
  }
  defer listener.Close()
  log.Println("Listening on port 8080")
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatal("Error accepting connection:", err)
    }
    log.Print(conn)
  }
}
