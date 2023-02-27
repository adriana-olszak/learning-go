package main

import (
	"fmt"
	"net/http"

	"go-getting-started/demo-web-server/controllers"
)

func main() {
	fmt.Println("Starting Web Server")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) // nill specifies that we will use the defult serve
	fmt.Println("Web Server Started")
}
