package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Maryam-Yumna/FCL-Website-Backend/src/routes"
)

func main() {
	fmt.Println("Hello")
	router := routes.Router()
	router.Path("/")
	http.Handle("/", router)
	log.Println("Server Started localhost :3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
