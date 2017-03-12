package main

import (
	"log"
	"net/http"

	"github.com/didiyudha/go-mongodb-restful/routes"
)

func main() {
	route := routes.NewRouter()
	log.Println("Server is up on :8080 port")
	log.Fatalln(http.ListenAndServe(":8080", route))
}
