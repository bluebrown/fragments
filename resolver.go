package main

import (
	"github.com/graph-gophers/graphql-go"
)

//
//
// ! Root Resolver

// Resolver is the root resolver.
type Resolver struct{}

//
// ** User **

//User asks the model politely.
func (*Resolver) User(args struct{ ID graphql.ID }) *User {
	return Model.UID(args)
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
