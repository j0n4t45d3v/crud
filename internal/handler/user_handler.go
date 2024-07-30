package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/crud/internal/entity"
	"github.com/j0n4t45d3v/crud/internal/repository"
)

type ResponseUser[T any] struct {
	Timestamp string `json:"timestamp"`
	Status    int16  `json:"status"`
	Data      T      `json:"data"`
}

type ResponseSucess struct {
	Timestamp string `json:"timestamp"`
	Status    int16  `json:"status"`
	Message   string `json:"message"`
}

type ResponseError struct {
	Timestamp string `json:"timestamp"`
	Status    int16  `json:"status"`
	Error     string `json:"error"`
}

type Controller struct {
	repository repository.IUserRepository
}

func NewController(repository repository.IUserRepository) Controller {
	return Controller{repository: repository}
}

func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.repository.FindAll()

	if err != nil {
		responseError := ResponseError{
			Timestamp: time.Now().String(),
			Status:    500,
			Error:     err.Error(),
		}
		fmt.Fprint(w, toJson(responseError))
		return
	}

	response := ResponseUser[[]entity.User]{
		Timestamp: time.Now().String(),
		Status:    200,
		Data:      users,
	}

	fmt.Fprint(w, toJson(response))
}

func (c Controller) Save(w http.ResponseWriter, r *http.Request) {
	var body entity.User
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		responseError := ResponseError{
			Timestamp: time.Now().String(),
			Status:    500,
			Error:     err.Error(),
		}
		fmt.Fprint(w, toJson(responseError))
		return
	}

	id, err := c.repository.Save(body)
	if err != nil {
		responseError := ResponseError{
			Timestamp: time.Now().String(),
			Status:    500,
			Error:     err.Error(),
		}
		fmt.Fprint(w, toJson(responseError))
		return
	}

	host := r.Host
	uriLocation := fmt.Sprintf("%v/v1/users/%d", host, id)
	w.Header().Add("Location", uriLocation)

	fmt.Fprint(w, toJson(body))
}

func (c Controller) Delete(w http.ResponseWriter, r *http.Request) {
	varsRequest := mux.Vars(r)
	id := varsRequest["id"]

	err := c.repository.Delete(id)
	if err != nil {
		responseError := ResponseError{
			Timestamp: time.Now().String(),
			Status:    500,
			Error:     err.Error(),
		}
		fmt.Fprint(w, toJson(responseError))
		return
	}

	res := ResponseSucess{
		Timestamp: time.Now().String(),
		Status:    200,
		Message:   "User Removed",
	}
	fmt.Fprint(w, toJson(res))
}

func (c Controller) Edit(w http.ResponseWriter, r *http.Request) {
	var body entity.User
	id := mux.Vars(r)["id"]
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		responseError := ResponseError{
			Timestamp: time.Now().String(),
			Status:    500,
			Error:     err.Error(),
		}
		fmt.Fprint(w, toJson(responseError))
		return
	}

	errUpdate := c.repository.Update(body, id)
	if errUpdate != nil {
		responseError := ResponseError{
			Timestamp: time.Now().String(),
			Status:    500,
			Error:     errUpdate.Error(),
		}
		fmt.Fprint(w, toJson(responseError))
		return
	}

	res := ResponseSucess{
		Timestamp: time.Now().String(),
		Status:    200,
		Message:   "User Edited",
	}
	fmt.Fprint(w, toJson(res))
}

func toJson[T any](r T) string {
	res, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return string(res)
}
