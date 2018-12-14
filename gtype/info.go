package gtype

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Info represents gql info type.
type Info struct {
	GID  graphql.ID
	GTxt string
}

// ID resolves info.id query.
func (i *Info) ID() graphql.ID { return i.GID }

// Text resolves info.text query.
func (i *Info) Text() string { return i.GTxt }

//
