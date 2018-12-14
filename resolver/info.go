package resolver

import (
	"github.com/bluebrown/fragments/gtype"
	graphql "github.com/graph-gophers/graphql-go"
)

// Info resolves info info query.
func (*Resolver) Info(args struct{ ID graphql.ID }) *gtype.Info {
	return mockInfo
}
