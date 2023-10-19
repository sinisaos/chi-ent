// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sinisaos/chi-ent/ent/predicate"
	"github.com/sinisaos/chi-ent/ent/question"
	"github.com/sinisaos/chi-ent/ent/questiontag"
	"github.com/sinisaos/chi-ent/ent/tag"
)

// QuestionTagUpdate is the builder for updating QuestionTag entities.
type QuestionTagUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionTagMutation
}

// Where appends a list predicates to the QuestionTagUpdate builder.
func (qtu *QuestionTagUpdate) Where(ps ...predicate.QuestionTag) *QuestionTagUpdate {
	qtu.mutation.Where(ps...)
	return qtu
}

// SetQuestionID sets the "question_id" field.
func (qtu *QuestionTagUpdate) SetQuestionID(i int) *QuestionTagUpdate {
	qtu.mutation.SetQuestionID(i)
	return qtu
}

// SetTagID sets the "tag_id" field.
func (qtu *QuestionTagUpdate) SetTagID(i int) *QuestionTagUpdate {
	qtu.mutation.SetTagID(i)
	return qtu
}

// SetQuestion sets the "question" edge to the Question entity.
func (qtu *QuestionTagUpdate) SetQuestion(q *Question) *QuestionTagUpdate {
	return qtu.SetQuestionID(q.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (qtu *QuestionTagUpdate) SetTag(t *Tag) *QuestionTagUpdate {
	return qtu.SetTagID(t.ID)
}

// Mutation returns the QuestionTagMutation object of the builder.
func (qtu *QuestionTagUpdate) Mutation() *QuestionTagMutation {
	return qtu.mutation
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qtu *QuestionTagUpdate) ClearQuestion() *QuestionTagUpdate {
	qtu.mutation.ClearQuestion()
	return qtu
}

// ClearTag clears the "tag" edge to the Tag entity.
func (qtu *QuestionTagUpdate) ClearTag() *QuestionTagUpdate {
	qtu.mutation.ClearTag()
	return qtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qtu *QuestionTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, qtu.sqlSave, qtu.mutation, qtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qtu *QuestionTagUpdate) SaveX(ctx context.Context) int {
	affected, err := qtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qtu *QuestionTagUpdate) Exec(ctx context.Context) error {
	_, err := qtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qtu *QuestionTagUpdate) ExecX(ctx context.Context) {
	if err := qtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qtu *QuestionTagUpdate) check() error {
	if _, ok := qtu.mutation.QuestionID(); qtu.mutation.QuestionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "QuestionTag.question"`)
	}
	if _, ok := qtu.mutation.TagID(); qtu.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "QuestionTag.tag"`)
	}
	return nil
}

func (qtu *QuestionTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := qtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(questiontag.Table, questiontag.Columns, sqlgraph.NewFieldSpec(questiontag.FieldID, field.TypeInt))
	if ps := qtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if qtu.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.QuestionTable,
			Columns: []string{questiontag.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qtu.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.QuestionTable,
			Columns: []string{questiontag.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qtu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.TagTable,
			Columns: []string{questiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qtu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.TagTable,
			Columns: []string{questiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questiontag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	qtu.mutation.done = true
	return n, nil
}

// QuestionTagUpdateOne is the builder for updating a single QuestionTag entity.
type QuestionTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionTagMutation
}

// SetQuestionID sets the "question_id" field.
func (qtuo *QuestionTagUpdateOne) SetQuestionID(i int) *QuestionTagUpdateOne {
	qtuo.mutation.SetQuestionID(i)
	return qtuo
}

// SetTagID sets the "tag_id" field.
func (qtuo *QuestionTagUpdateOne) SetTagID(i int) *QuestionTagUpdateOne {
	qtuo.mutation.SetTagID(i)
	return qtuo
}

// SetQuestion sets the "question" edge to the Question entity.
func (qtuo *QuestionTagUpdateOne) SetQuestion(q *Question) *QuestionTagUpdateOne {
	return qtuo.SetQuestionID(q.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (qtuo *QuestionTagUpdateOne) SetTag(t *Tag) *QuestionTagUpdateOne {
	return qtuo.SetTagID(t.ID)
}

// Mutation returns the QuestionTagMutation object of the builder.
func (qtuo *QuestionTagUpdateOne) Mutation() *QuestionTagMutation {
	return qtuo.mutation
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qtuo *QuestionTagUpdateOne) ClearQuestion() *QuestionTagUpdateOne {
	qtuo.mutation.ClearQuestion()
	return qtuo
}

// ClearTag clears the "tag" edge to the Tag entity.
func (qtuo *QuestionTagUpdateOne) ClearTag() *QuestionTagUpdateOne {
	qtuo.mutation.ClearTag()
	return qtuo
}

// Where appends a list predicates to the QuestionTagUpdate builder.
func (qtuo *QuestionTagUpdateOne) Where(ps ...predicate.QuestionTag) *QuestionTagUpdateOne {
	qtuo.mutation.Where(ps...)
	return qtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (qtuo *QuestionTagUpdateOne) Select(field string, fields ...string) *QuestionTagUpdateOne {
	qtuo.fields = append([]string{field}, fields...)
	return qtuo
}

// Save executes the query and returns the updated QuestionTag entity.
func (qtuo *QuestionTagUpdateOne) Save(ctx context.Context) (*QuestionTag, error) {
	return withHooks(ctx, qtuo.sqlSave, qtuo.mutation, qtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qtuo *QuestionTagUpdateOne) SaveX(ctx context.Context) *QuestionTag {
	node, err := qtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (qtuo *QuestionTagUpdateOne) Exec(ctx context.Context) error {
	_, err := qtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qtuo *QuestionTagUpdateOne) ExecX(ctx context.Context) {
	if err := qtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qtuo *QuestionTagUpdateOne) check() error {
	if _, ok := qtuo.mutation.QuestionID(); qtuo.mutation.QuestionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "QuestionTag.question"`)
	}
	if _, ok := qtuo.mutation.TagID(); qtuo.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "QuestionTag.tag"`)
	}
	return nil
}

func (qtuo *QuestionTagUpdateOne) sqlSave(ctx context.Context) (_node *QuestionTag, err error) {
	if err := qtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(questiontag.Table, questiontag.Columns, sqlgraph.NewFieldSpec(questiontag.FieldID, field.TypeInt))
	id, ok := qtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "QuestionTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := qtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, questiontag.FieldID)
		for _, f := range fields {
			if !questiontag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != questiontag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := qtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if qtuo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.QuestionTable,
			Columns: []string{questiontag.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qtuo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.QuestionTable,
			Columns: []string{questiontag.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qtuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.TagTable,
			Columns: []string{questiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qtuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   questiontag.TagTable,
			Columns: []string{questiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &QuestionTag{config: qtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, qtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questiontag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	qtuo.mutation.done = true
	return _node, nil
}
