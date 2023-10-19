// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sinisaos/chi-ent/ent/predicate"
	"github.com/sinisaos/chi-ent/ent/questiontag"
)

// QuestionTagDelete is the builder for deleting a QuestionTag entity.
type QuestionTagDelete struct {
	config
	hooks    []Hook
	mutation *QuestionTagMutation
}

// Where appends a list predicates to the QuestionTagDelete builder.
func (qtd *QuestionTagDelete) Where(ps ...predicate.QuestionTag) *QuestionTagDelete {
	qtd.mutation.Where(ps...)
	return qtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qtd *QuestionTagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, qtd.sqlExec, qtd.mutation, qtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (qtd *QuestionTagDelete) ExecX(ctx context.Context) int {
	n, err := qtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qtd *QuestionTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(questiontag.Table, sqlgraph.NewFieldSpec(questiontag.FieldID, field.TypeInt))
	if ps := qtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, qtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	qtd.mutation.done = true
	return affected, err
}

// QuestionTagDeleteOne is the builder for deleting a single QuestionTag entity.
type QuestionTagDeleteOne struct {
	qtd *QuestionTagDelete
}

// Where appends a list predicates to the QuestionTagDelete builder.
func (qtdo *QuestionTagDeleteOne) Where(ps ...predicate.QuestionTag) *QuestionTagDeleteOne {
	qtdo.qtd.mutation.Where(ps...)
	return qtdo
}

// Exec executes the deletion query.
func (qtdo *QuestionTagDeleteOne) Exec(ctx context.Context) error {
	n, err := qtdo.qtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{questiontag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qtdo *QuestionTagDeleteOne) ExecX(ctx context.Context) {
	if err := qtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
