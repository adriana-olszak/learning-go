package models // pacakge name equals the directory name

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// variable block
var (
	users  []*User // slice of User pointers, by working with pointers we are gonna be able to manipulate the user from various places in the app
	nextID = 1     // compiler will specify the data type by itself, we dont need to use colon inside of variable block
)

func GetUsers() []*User { // return slice of User pointers
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("User cannot include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u) // append another pointer to new user
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with id `%v` not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, oldU := range users {
		if oldU.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id `%v` not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			// remove from slice of pointers, basically doing a splice
			users = append(
				users[:i],      // get users before the user that feound
				users[i+1:]..., // get users after the user that we found
			)
			return nil
		}
	}

	return fmt.Errorf("User with id `%v` not found", id)
}
