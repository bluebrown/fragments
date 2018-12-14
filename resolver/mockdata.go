package resolver

import "github.com/bluebrown/fragments/gtype"

var mockInfo = &gtype.Info{
	GID:  "some id",
	GTxt: "Hello, GraphQL!",
}

var mockUser1 = &gtype.User{
	GID:   "Some id",
	GName: "Jim",
	GMail: "gmail",
}
var mockUser2 = &gtype.User{
	GID:   "Some id",
	GName: "frank",
	GMail: "hotmail",
}

var mockMessage1 = &gtype.Message{
	GID:   "1",
	GText: "hi there",
	GFrom: mockUser1,
	GTo:   []*gtype.User{mockUser1, mockUser2},
}

var mockMessage2 = &gtype.Message{
	GID:   "2",
	GText: "yo mac",
	GFrom: mockUser2,
	GTo:   []*gtype.User{mockUser2, mockUser1},
}
