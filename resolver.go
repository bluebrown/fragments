package main

//
//
//** Root Resolver **

// Resolver is the root resolver.
type Resolver struct{}

// Hello is a simple example-query.
func (*Resolver) Hello() string {
	return "Hello, world!"
}

//User resolves the user query.
func (*Resolver) User() *User {
	return store.popUser()
}

//Users resolves the users query.
func (*Resolver) Users() []*User {
	return store.popUsers()
}

//
//
//** Type User **

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
