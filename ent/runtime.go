// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/sinisaos/chi-ent/ent/answer"
	"github.com/sinisaos/chi-ent/ent/question"
	"github.com/sinisaos/chi-ent/ent/schema"
	"github.com/sinisaos/chi-ent/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescCreatedAt is the schema descriptor for created_at field.
	answerDescCreatedAt := answerFields[1].Descriptor()
	// answer.DefaultCreatedAt holds the default value on creation for the created_at field.
	answer.DefaultCreatedAt = answerDescCreatedAt.Default.(func() time.Time)
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescCreatedAt is the schema descriptor for created_at field.
	questionDescCreatedAt := questionFields[2].Descriptor()
	// question.DefaultCreatedAt holds the default value on creation for the created_at field.
	question.DefaultCreatedAt = questionDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
