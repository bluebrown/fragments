package main

/* Simple mock data struct that
   provides functionality to
   handle query's from graphql-types. */
type db struct {
	users []User
}

// Return a single user.
func (_db *db) popUser() *User {
	return &_db.users[0]
}

// Return userlist.
func (_db *db) popUsers() []*User {
	u := []*User{}
	for key := range _db.users {
		u = append(u, &_db.users[key])
	}
	return u
}

// Store singleton stores mock data
// and responds to graphql-type-requests.
var store = db{
	users: []User{
		User{
			name: "John Coyle",
			mail: "coyle@gmail.com",
		},
		User{
			name: "Frank Jones",
			mail: "slaptstick@montecarl.io",
		},
	},
}
