package main

import (
	"net/http"

	"go-getting-started/demo-web-server/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) // nill specifies that we will use the defult serve
}
