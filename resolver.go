package main

//** Root Resolver **

// Resolver is the root resolver.
type Resolver struct{}

// Hello is a simple example-query.
func (*Resolver) Hello() string {
	return "Hello, world!"
}

//User resolves the user query.
func (*Resolver) User() *User {
	return &store.users[0]
}

//Users resolves the users query.
func (*Resolver) Users() []*User {
	u := []*User{}
	for key := range store.users {
		u = append(u, &store.users[key])
	}
	return u
}

func (*Resolver) Message() *Message {
	return &store.messages[0]
}

//
//** Type User **

// User implements the user type.
type User struct {
	name string
	mail string
}

// Name resolves the User-Name query.
func (u *User) Name() string {
	return u.name
}

// Mail resolves the User-Mail query.
func (u *User) Mail() string {
	return u.mail
}

//
//** Type Message **

type Message struct {
	text string
	from *User
	to   []*User
}

func (m *Message) Text() string {
	return m.text
}
func (m *Message) From() *User {
	return m.from
}
func (m *Message) To() []*User {
	return m.to
}
