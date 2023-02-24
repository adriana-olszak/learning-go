package controllers

import (
	"net/http"
	"regexp"
)

// it is common to have an empty struct to keep methods on it.
type userController struct {
	userIDPattern *regexp.Regexp
}

// method definition, binds a function to a specific type
// there is no Constructor in go
// if our function name, in this case ServeHTTP, mathces interface defintion it will be used automatically
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the USer Controller!"))
}

func newUserController() *userController {
	return &userController{ // we can immidietly return memory address of a struct initialisation, go will automatically promote localvariable address to public value
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
