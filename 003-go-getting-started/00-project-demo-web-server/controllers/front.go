package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	// we need to have users route with and without ending slash provided
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// again we need to manually specificy that we need to decode Go struct into JSON
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
