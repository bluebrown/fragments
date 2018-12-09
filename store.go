package main

// Simple mock data struct that provides functionality to
// handle query's from graphql-types.
type db struct {
	users    []User
	messages []Message
}

// ! Singleton that stores mock data
// and responds to graphql-type-requests.
var store = db{}

func (_db *db) init() {

	_db.users = []User{
		User{
			id:   "1",
			name: "John Coyle",
			mail: "coyle@gmail.com",
		},
		User{
			id:   "2",
			name: "Frank Jones",
			mail: "slaptstick@montecarl.io",
		},
		User{
			id:   "3",
			name: "Hank Arnold",
			mail: "macbane@outfox.pong",
		},
	}
	_db.messages = []Message{
		Message{
			text: "Hello, GraphQL!",
			from: &_db.users[0],
			to: []*User{
				&_db.users[1],
				&_db.users[2],
			},
		},
		Message{
			text: "Thanks :)",
			from: &_db.users[1],
			to: []*User{
				&_db.users[0],
			},
		},
	}
}
