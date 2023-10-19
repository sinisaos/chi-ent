package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type QuestionTag struct {
	ent.Schema
}

// Fields of the QuestionTag.
func (QuestionTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("question_id"),
		field.Int("tag_id"),
	}
}

// Edges of the QuestionTag.
func (QuestionTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("question", Question.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Unique().
			Required().
			Field("question_id"),
		edge.To("tag", Tag.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Unique().
			Required().
			Field("tag_id"),
	}
}

// Indexes of the QuestionTag.
func (QuestionTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag_id").
			Unique(),
	}
}
