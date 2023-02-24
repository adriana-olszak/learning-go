package main

import (
	"fmt"

	"go-getting-started/demo-web-server/models"
)

func main() {
	user := models.User{
		ID:        2,
		FirstName: "Marc",
		LastName:  "Loleluis",
	}

	fmt.Println(user)
}
