package main

import graphql "github.com/graph-gophers/graphql-go"

//
// ! Type User

// User implements the user type.
type User struct {
	id   graphql.ID
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

// ID resolves the User-ID query.
func (u *User) ID() graphql.ID {
	return u.id
}

//
// ! Type Message

// Message implents Message type.
type Message struct {
	text string
	from *User
	to   []*User
}

// Text returns message text.
func (m *Message) Text() string {
	return m.text
}

// From returns pointer to message-senders
// memory address.
func (m *Message) From() *User {
	return m.from
}

// To return array of pointer to message-senders
// memory adress.
func (m *Message) To() []*User {
	return m.to
}
