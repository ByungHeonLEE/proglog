package main

import (
	"log"

	"github.com/ByungHeonLEE/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
