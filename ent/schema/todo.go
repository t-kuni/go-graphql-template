package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Bool("done"),
		field.Int("user_id"),
		field.Time("created_at").Default(time.Now),
	}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("todos").Unique().Field("user_id").Required(),
	}
}
