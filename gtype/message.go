package gtype

// Message implents Message type.
type Message struct {
	GText string
	GFrom *User
	GTo   []*User
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
