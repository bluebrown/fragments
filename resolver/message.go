package resolver

import "github.com/bluebrown/fragments/gtype"

// Message resolves the message query.
func (*Resolver) Message() *gtype.Message {
	return mockMessage1
}

// Messages resolves message list query.
func (*Resolver) Messages() []*gtype.Message {
	return []*gtype.Message{mockMessage1, mockMessage2}
}
