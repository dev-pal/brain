package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dev-pal/brain/pkg/api/rest"
)

func main() {
	router := rest.Handler()

	fmt.Println("The dev-pal brain server is on tap now: http://localhost:64800")
	log.Fatal(http.ListenAndServe(":64800", router))
}
