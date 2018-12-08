package main

// Simple mock database that implements gql types.
type db struct {
	users []User
}

func (_db *db) popUser() *User {
	return &_db.users[0]
}

var store = db{
	users: []User{
		User{name: "Tony Test"},
	},
}
