package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"time"
)

func withBase(fields []ent.Field) []ent.Field {
	return append([]ent.Field{
		field.Uint("id").
			StructTag(`json:"oid,omitempty"`),
		field.Time("created_at").
			Default(time.Now),
	}, fields...)
}
