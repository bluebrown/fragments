package main

//
//
// ! Root Resolver

// Resolver is the root resolver.
type Resolver struct{}

//
// ** User **

// User resolves the user query.
func (*Resolver) User(args struct{ Name string }) *User {
	for key, val := range store.users {
		if args.Name == val.name {
			return &store.users[key]
		}
	}
	return &User{"None", "None"}
}

//Users resolves the users query.
func (*Resolver) Users() []*User {
	u := []*User{}
	for key := range store.users {
		u = append(u, &store.users[key])
	}
	return u
}

//
// ** Message **

// Message resolves the message query.
func (*Resolver) Message() *Message {
	return &store.messages[0]
}

// Messages resolves message list query.
func (*Resolver) Messages() []*Message {
	m := []*Message{}
	for key := range store.messages {
		m = append(m, &store.messages[key])
	}
	return m
}

//
//
// ! Type User

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
