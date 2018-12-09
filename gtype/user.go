package gtype

import graphql "github.com/graph-gophers/graphql-go"

// User implements the user type.
type User struct {
	GID   graphql.ID
	GName string
	GMail string
}

// Name resolves the User-Name query.
func (u *User) Name() string {
	return u.GName
}

// Mail resolves the User-Mail query.
func (u *User) Mail() string {
	return u.GMail
}

// ID resolves the User-ID query.
func (u *User) ID() graphql.ID {
	return u.GID
}

//User asks the model politely.
func (*Resolver) User(args struct{ ID graphql.ID }) *User {
	return &User{args.ID, "Jim", "gmail"}
}

//Users resolves the users query.
func (*Resolver) Users() []*User {
	return []*User{
		&User{"2", "Hans", "pmail"},
		&User{"3", "john", "jmail"},
	}
}
