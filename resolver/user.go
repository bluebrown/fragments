package resolver

import (
	"github.com/bluebrown/fragments/gtype"
	graphql "github.com/graph-gophers/graphql-go"
)

//User resolves the user query.
func (*Resolver) User(args struct{ ID graphql.ID }) *gtype.User {
	return mockUser1
}

//Users resolves the users query.
func (*Resolver) Users() []*gtype.User {
	return []*gtype.User{mockUser1, mockUser2}
}
