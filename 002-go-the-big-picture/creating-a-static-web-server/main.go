package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// functions are first class citiziens in Go, you can create it on the fly
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string

		if len(names) == 1 {
			name = names[0]
		}

		m := map[string]string{"name": name} // simple map of string to string in TS Map<string, string>, with first value of {name:name}

		// we need to encode this map into JSON, we need to specifiy some kind of writer to write JSON into
		enc := json.NewEncoder(w)

		// encoding and writng to the passed writer 
		enc.Encode(m)
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

// [PROBLEMS]
// [1]
// go build .
// # go-the-big-picture/creating-a-static-web-server
// ./main.go:10:18: cannot use '/' (untyped rune constant 47) as string value in argument to http.HandleFunc
// AGAIN DUE TO VSC NOT SAVED. Strangely enough I did set it to save on window loosing focus.
// Looks like Hotkey Warp terminal does not cause it loos focus?
// Setting it up so it always saves.
