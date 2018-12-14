package resolver

import (
	"github.com/bluebrown/fragments/gtype"
	graphql "github.com/graph-gophers/graphql-go"
)

// Resolver is the root resolver and can
// peform all the resolving function
// for graph-queries.
type Resolver struct{}

// Info resolves info info query.
func (*Resolver) Info(args struct{ ID graphql.ID }) *gtype.Info {
	return &gtype.Info{
		GID:  args.ID,
		GTxt: "Hello, GraphQL!",
	}
}

//User resolves the user query.
func (*Resolver) User(args struct{ ID graphql.ID }) *gtype.User {
	return &gtype.User{args.ID, "Jim", "gmail"}
}

//Users resolves the users query.
func (*Resolver) Users() []*gtype.User {
	return []*gtype.User{
		&gtype.User{"2", "Hans", "pmail"},
		&gtype.User{"3", "john", "jmail"},
	}
}

//
//* Resolver =>

// Message resolves the message query.
func (*Resolver) Message() *gtype.Message {
	return &gtype.Message{"1", &gtype.User{}, []*gtype.User{}}
}

// Messages resolves message list query.
func (*Resolver) Messages() []*gtype.Message {
	return []*gtype.Message{
		&gtype.Message{"2", &gtype.User{}, []*gtype.User{}},
		&gtype.Message{"3", &gtype.User{}, []*gtype.User{}},
	}
}
