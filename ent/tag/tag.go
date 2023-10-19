// Code generated by ent, DO NOT EDIT.

package tag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeQuestions holds the string denoting the questions edge name in mutations.
	EdgeQuestions = "questions"
	// EdgeTagQuestion holds the string denoting the tag_question edge name in mutations.
	EdgeTagQuestion = "tag_question"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// QuestionsTable is the table that holds the questions relation/edge. The primary key declared below.
	QuestionsTable = "question_tags"
	// QuestionsInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionsInverseTable = "questions"
	// TagQuestionTable is the table that holds the tag_question relation/edge.
	TagQuestionTable = "question_tags"
	// TagQuestionInverseTable is the table name for the QuestionTag entity.
	// It exists in this package in order to avoid circular dependency with the "questiontag" package.
	TagQuestionInverseTable = "question_tags"
	// TagQuestionColumn is the table column denoting the tag_question relation/edge.
	TagQuestionColumn = "tag_id"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tags"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_tags",
}

var (
	// QuestionsPrimaryKey and QuestionsColumn2 are the table columns denoting the
	// primary key for the questions relation (M2M).
	QuestionsPrimaryKey = []string{"question_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByQuestionsCount orders the results by questions count.
func ByQuestionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newQuestionsStep(), opts...)
	}
}

// ByQuestions orders the results by questions terms.
func ByQuestions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagQuestionCount orders the results by tag_question count.
func ByTagQuestionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagQuestionStep(), opts...)
	}
}

// ByTagQuestion orders the results by tag_question terms.
func ByTagQuestion(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagQuestionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newQuestionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, QuestionsTable, QuestionsPrimaryKey...),
	)
}
func newTagQuestionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagQuestionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, TagQuestionTable, TagQuestionColumn),
	)
}
