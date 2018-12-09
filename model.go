package main

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Simple model to interact with mock store
type model struct{}

func (model) UserByID(ID graphql.ID) *User {
	for key, val := range store.users {
		if ID == val.id {
			return &store.users[key]
		}
	}
	return &User{"-1", "None", "None"}
}

// Modeler interface for flexible driver use
type modeler interface {
	UserByID(ID graphql.ID) *User
}

// driver that holds the model and some configurations.
type driver struct{ model modeler }

func (d *driver) UID(args struct{ ID graphql.ID }) *User {
	return d.model.UserByID(args.ID)
}
