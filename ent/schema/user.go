package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return withBase([]ent.Field{
		field.String("name").
			Default("unknown"),
		field.String("email").
			Unique(),
		field.String("password").
			Default("unknown"),
	})
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
