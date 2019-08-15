package main

import (
	"fmt"
	"log"
	"net/http"

	"go-bioreg/server/routes"
)

func main() {
	PORT := "3000"
	fmt.Println("Serving on port: ", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, routes.MyRouter()))
}