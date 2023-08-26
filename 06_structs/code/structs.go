package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group represents a set of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func describeGroup(g Group) string {
	// "This user group has 19 users. The newest user is Ibrahim. Accepting New Users: "
	if len(g.users) >= 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting New Users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 1, FirstName: "Christopher", LastName: "Nollan", Email: "christopher.nolan@gmail.com"}
	g := Group{
		role:           "admin",
		users:          []User{u1, u2},
		newestUser:     u2,
		spaceAvailable: true,
	}

	g.spaceAvailable = false
	fmt.Println(describeUser(u1))
	fmt.Println(describeGroup(g))
	fmt.Println(g)
}
