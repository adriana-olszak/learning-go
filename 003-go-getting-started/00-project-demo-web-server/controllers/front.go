package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	// we need to have users route with and without ending slash provided
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
