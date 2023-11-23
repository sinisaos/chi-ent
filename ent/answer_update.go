// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sinisaos/chi-ent/ent/answer"
	"github.com/sinisaos/chi-ent/ent/predicate"
	"github.com/sinisaos/chi-ent/ent/question"
	"github.com/sinisaos/chi-ent/ent/user"
)

// AnswerUpdate is the builder for updating Answer entities.
type AnswerUpdate struct {
	config
	hooks    []Hook
	mutation *AnswerMutation
}

// Where appends a list predicates to the AnswerUpdate builder.
func (au *AnswerUpdate) Where(ps ...predicate.Answer) *AnswerUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetContent sets the "content" field.
func (au *AnswerUpdate) SetContent(s string) *AnswerUpdate {
	au.mutation.SetContent(s)
	return au
}

// SetLikes sets the "likes" field.
func (au *AnswerUpdate) SetLikes(i int) *AnswerUpdate {
	au.mutation.ResetLikes()
	au.mutation.SetLikes(i)
	return au
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (au *AnswerUpdate) SetNillableLikes(i *int) *AnswerUpdate {
	if i != nil {
		au.SetLikes(*i)
	}
	return au
}

// AddLikes adds i to the "likes" field.
func (au *AnswerUpdate) AddLikes(i int) *AnswerUpdate {
	au.mutation.AddLikes(i)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AnswerUpdate) SetUpdatedAt(t time.Time) *AnswerUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetIsAcceptedAnswer sets the "is_accepted_answer" field.
func (au *AnswerUpdate) SetIsAcceptedAnswer(b bool) *AnswerUpdate {
	au.mutation.SetIsAcceptedAnswer(b)
	return au
}

// SetNillableIsAcceptedAnswer sets the "is_accepted_answer" field if the given value is not nil.
func (au *AnswerUpdate) SetNillableIsAcceptedAnswer(b *bool) *AnswerUpdate {
	if b != nil {
		au.SetIsAcceptedAnswer(*b)
	}
	return au
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (au *AnswerUpdate) SetQuestionID(id int) *AnswerUpdate {
	au.mutation.SetQuestionID(id)
	return au
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (au *AnswerUpdate) SetNillableQuestionID(id *int) *AnswerUpdate {
	if id != nil {
		au = au.SetQuestionID(*id)
	}
	return au
}

// SetQuestion sets the "question" edge to the Question entity.
func (au *AnswerUpdate) SetQuestion(q *Question) *AnswerUpdate {
	return au.SetQuestionID(q.ID)
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (au *AnswerUpdate) SetAuthorID(id int) *AnswerUpdate {
	au.mutation.SetAuthorID(id)
	return au
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (au *AnswerUpdate) SetNillableAuthorID(id *int) *AnswerUpdate {
	if id != nil {
		au = au.SetAuthorID(*id)
	}
	return au
}

// SetAuthor sets the "author" edge to the User entity.
func (au *AnswerUpdate) SetAuthor(u *User) *AnswerUpdate {
	return au.SetAuthorID(u.ID)
}

// Mutation returns the AnswerMutation object of the builder.
func (au *AnswerUpdate) Mutation() *AnswerMutation {
	return au.mutation
}

// ClearQuestion clears the "question" edge to the Question entity.
func (au *AnswerUpdate) ClearQuestion() *AnswerUpdate {
	au.mutation.ClearQuestion()
	return au
}

// ClearAuthor clears the "author" edge to the User entity.
func (au *AnswerUpdate) ClearAuthor() *AnswerUpdate {
	au.mutation.ClearAuthor()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AnswerUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnswerUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnswerUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnswerUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AnswerUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := answer.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AnswerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(answer.Table, answer.Columns, sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Content(); ok {
		_spec.SetField(answer.FieldContent, field.TypeString, value)
	}
	if value, ok := au.mutation.Likes(); ok {
		_spec.SetField(answer.FieldLikes, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedLikes(); ok {
		_spec.AddField(answer.FieldLikes, field.TypeInt, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(answer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.IsAcceptedAnswer(); ok {
		_spec.SetField(answer.FieldIsAcceptedAnswer, field.TypeBool, value)
	}
	if au.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
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
	if au.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.AuthorTable,
			Columns: []string{answer.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.AuthorTable,
			Columns: []string{answer.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{answer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AnswerUpdateOne is the builder for updating a single Answer entity.
type AnswerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AnswerMutation
}

// SetContent sets the "content" field.
func (auo *AnswerUpdateOne) SetContent(s string) *AnswerUpdateOne {
	auo.mutation.SetContent(s)
	return auo
}

// SetLikes sets the "likes" field.
func (auo *AnswerUpdateOne) SetLikes(i int) *AnswerUpdateOne {
	auo.mutation.ResetLikes()
	auo.mutation.SetLikes(i)
	return auo
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (auo *AnswerUpdateOne) SetNillableLikes(i *int) *AnswerUpdateOne {
	if i != nil {
		auo.SetLikes(*i)
	}
	return auo
}

// AddLikes adds i to the "likes" field.
func (auo *AnswerUpdateOne) AddLikes(i int) *AnswerUpdateOne {
	auo.mutation.AddLikes(i)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AnswerUpdateOne) SetUpdatedAt(t time.Time) *AnswerUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetIsAcceptedAnswer sets the "is_accepted_answer" field.
func (auo *AnswerUpdateOne) SetIsAcceptedAnswer(b bool) *AnswerUpdateOne {
	auo.mutation.SetIsAcceptedAnswer(b)
	return auo
}

// SetNillableIsAcceptedAnswer sets the "is_accepted_answer" field if the given value is not nil.
func (auo *AnswerUpdateOne) SetNillableIsAcceptedAnswer(b *bool) *AnswerUpdateOne {
	if b != nil {
		auo.SetIsAcceptedAnswer(*b)
	}
	return auo
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (auo *AnswerUpdateOne) SetQuestionID(id int) *AnswerUpdateOne {
	auo.mutation.SetQuestionID(id)
	return auo
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (auo *AnswerUpdateOne) SetNillableQuestionID(id *int) *AnswerUpdateOne {
	if id != nil {
		auo = auo.SetQuestionID(*id)
	}
	return auo
}

// SetQuestion sets the "question" edge to the Question entity.
func (auo *AnswerUpdateOne) SetQuestion(q *Question) *AnswerUpdateOne {
	return auo.SetQuestionID(q.ID)
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (auo *AnswerUpdateOne) SetAuthorID(id int) *AnswerUpdateOne {
	auo.mutation.SetAuthorID(id)
	return auo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (auo *AnswerUpdateOne) SetNillableAuthorID(id *int) *AnswerUpdateOne {
	if id != nil {
		auo = auo.SetAuthorID(*id)
	}
	return auo
}

// SetAuthor sets the "author" edge to the User entity.
func (auo *AnswerUpdateOne) SetAuthor(u *User) *AnswerUpdateOne {
	return auo.SetAuthorID(u.ID)
}

// Mutation returns the AnswerMutation object of the builder.
func (auo *AnswerUpdateOne) Mutation() *AnswerMutation {
	return auo.mutation
}

// ClearQuestion clears the "question" edge to the Question entity.
func (auo *AnswerUpdateOne) ClearQuestion() *AnswerUpdateOne {
	auo.mutation.ClearQuestion()
	return auo
}

// ClearAuthor clears the "author" edge to the User entity.
func (auo *AnswerUpdateOne) ClearAuthor() *AnswerUpdateOne {
	auo.mutation.ClearAuthor()
	return auo
}

// Where appends a list predicates to the AnswerUpdate builder.
func (auo *AnswerUpdateOne) Where(ps ...predicate.Answer) *AnswerUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AnswerUpdateOne) Select(field string, fields ...string) *AnswerUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Answer entity.
func (auo *AnswerUpdateOne) Save(ctx context.Context) (*Answer, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnswerUpdateOne) SaveX(ctx context.Context) *Answer {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AnswerUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnswerUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AnswerUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := answer.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AnswerUpdateOne) sqlSave(ctx context.Context) (_node *Answer, err error) {
	_spec := sqlgraph.NewUpdateSpec(answer.Table, answer.Columns, sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Answer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, answer.FieldID)
		for _, f := range fields {
			if !answer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != answer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Content(); ok {
		_spec.SetField(answer.FieldContent, field.TypeString, value)
	}
	if value, ok := auo.mutation.Likes(); ok {
		_spec.SetField(answer.FieldLikes, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedLikes(); ok {
		_spec.AddField(answer.FieldLikes, field.TypeInt, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(answer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.IsAcceptedAnswer(); ok {
		_spec.SetField(answer.FieldIsAcceptedAnswer, field.TypeBool, value)
	}
	if auo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
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
	if auo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.AuthorTable,
			Columns: []string{answer.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.AuthorTable,
			Columns: []string{answer.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Answer{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{answer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
