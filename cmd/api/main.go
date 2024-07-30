package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/crud/internal/route"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	route.SetRoutesV1(r)
	fmt.Println("Running server in port: 8080")
	http.ListenAndServe(":8000", r)
}
