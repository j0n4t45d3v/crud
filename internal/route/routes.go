package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/crud/database"
	"github.com/j0n4t45d3v/crud/internal/handler"
	"github.com/j0n4t45d3v/crud/internal/repository"
)

func SetRoutesV1(r *mux.Router) {
	con := database.GetConnection()
	userRepo := repository.NewRepository(con)
	c := handler.NewController(userRepo)
	r.HandleFunc("/v1", homeRoute).Methods("GET")
	r.HandleFunc("/v1/users", c.GetAll).Methods("GET")
	r.HandleFunc("/v1/users", c.Save).Methods("POST")
	r.HandleFunc("/v1/users/{id}", c.Delete).Methods("DELETE")
	r.HandleFunc("/v1/users/{id}", c.Edit).Methods("PUT")
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message":   "Index Route",
		"timestamp": time.Now().String(),
		"status":    "200",
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(responseJson))
}
