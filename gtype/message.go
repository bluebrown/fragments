package gtype

import graphql "github.com/graph-gophers/graphql-go"

// Message implents Message type.
type Message struct {
	GID   graphql.ID
	GText string
	GFrom *User
	GTo   []*User
}

// ID resolves the ID query
func (m *Message) ID() graphql.ID {
	return m.GID
}

// Text returns message text.
func (m *Message) Text() string {
	return m.GText
}

// From returns pointer to message-senders
// memory address.
func (m *Message) From() *User {
	return m.GFrom
}

// To return array of pointer to message-senders
// memory adress.
func (m *Message) To() []*User {
	return m.GTo
}
