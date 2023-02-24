package models // pacakge name equals the directory name

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
	u.ID = nextID
	nextID++
	users = append(users, &u) // append another pointer to new user
	return u, nil
}
