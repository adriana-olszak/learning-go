package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"go-getting-started/demo-web-server/models"
)

// it is common to have an empty struct to keep methods on it.
type userController struct {
	userIDPattern *regexp.Regexp
}

// method definition, binds a function to a specific type
// there is no Constructor in go
// if our function name, in this case ServeHTTP, mathces interface defintion it will be used automatically
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		// get matching parameter
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)

		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		// parseInt on the parameter
		id, err := strconv.Atoi(matches[1])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.getOne(id, w)
		case http.MethodPut:
			uc.put(w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJSON(users, w)
}
func (uc userController) getOne(id int, w http.ResponseWriter) {
	user, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJSON(user, w)
}

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		fmt.Printf("Some error when parsing", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users, err := models.AddUser(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	encodeResponseAsJSON(users, w)
}

func (uc userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc userController) put(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedUser, err := models.UpdateUser(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJSON(updatedUser, w)

}

// constructor function
// as in Go there is no concept of constructor for Objects we can use function to create an instance with specified initial values
func newUserController() *userController {
	return &userController{ // we can immidietly return memory address of a struct initialisation, go will automatically promote localvariable address to public value
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

// we need to decode from JSON into Go type of choice
func (uc userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body) // define a decoder

	var u models.User
	err := dec.Decode(&u) // use decoder to decode into specified struct
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
