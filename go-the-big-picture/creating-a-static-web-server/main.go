package main

import (
	"log"
	"net/http"
)

func main() {
	// functions are first class citiziens in Go, you can create it on the fly
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World")) // []byte is needed bc writies always work with slices of bytes. We convert String into Byte using []byte
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
