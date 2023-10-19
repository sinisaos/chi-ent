// Code generated by ent, DO NOT EDIT.

package questiontag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the questiontag type in the database.
	Label = "question_tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuestionID holds the string denoting the question_id field in the database.
	FieldQuestionID = "question_id"
	// FieldTagID holds the string denoting the tag_id field in the database.
	FieldTagID = "tag_id"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// Table holds the table name of the questiontag in the database.
	Table = "question_tags"
	// QuestionTable is the table that holds the question relation/edge.
	QuestionTable = "question_tags"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "question_id"
	// TagTable is the table that holds the tag relation/edge.
	TagTable = "question_tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "tag_id"
)

// Columns holds all SQL columns for questiontag fields.
var Columns = []string{
	FieldID,
	FieldQuestionID,
	FieldTagID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the QuestionTag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuestionID orders the results by the question_id field.
func ByQuestionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuestionID, opts...).ToFunc()
}

// ByTagID orders the results by the tag_id field.
func ByTagID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTagID, opts...).ToFunc()
}

// ByQuestionField orders the results by question field.
func ByQuestionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionStep(), sql.OrderByField(field, opts...))
	}
}

// ByTagField orders the results by tag field.
func ByTagField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagStep(), sql.OrderByField(field, opts...))
	}
}
func newQuestionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, QuestionTable, QuestionColumn),
	)
}
func newTagStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
	)
}
