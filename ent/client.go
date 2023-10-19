// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/sinisaos/chi-ent/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sinisaos/chi-ent/ent/answer"
	"github.com/sinisaos/chi-ent/ent/question"
	"github.com/sinisaos/chi-ent/ent/questiontag"
	"github.com/sinisaos/chi-ent/ent/tag"
	"github.com/sinisaos/chi-ent/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Answer is the client for interacting with the Answer builders.
	Answer *AnswerClient
	// Question is the client for interacting with the Question builders.
	Question *QuestionClient
	// QuestionTag is the client for interacting with the QuestionTag builders.
	QuestionTag *QuestionTagClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Answer = NewAnswerClient(c.config)
	c.Question = NewQuestionClient(c.config)
	c.QuestionTag = NewQuestionTagClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Answer:      NewAnswerClient(cfg),
		Question:    NewQuestionClient(cfg),
		QuestionTag: NewQuestionTagClient(cfg),
		Tag:         NewTagClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Answer:      NewAnswerClient(cfg),
		Question:    NewQuestionClient(cfg),
		QuestionTag: NewQuestionTagClient(cfg),
		Tag:         NewTagClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Answer.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Answer.Use(hooks...)
	c.Question.Use(hooks...)
	c.QuestionTag.Use(hooks...)
	c.Tag.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Answer.Intercept(interceptors...)
	c.Question.Intercept(interceptors...)
	c.QuestionTag.Intercept(interceptors...)
	c.Tag.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AnswerMutation:
		return c.Answer.mutate(ctx, m)
	case *QuestionMutation:
		return c.Question.mutate(ctx, m)
	case *QuestionTagMutation:
		return c.QuestionTag.mutate(ctx, m)
	case *TagMutation:
		return c.Tag.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AnswerClient is a client for the Answer schema.
type AnswerClient struct {
	config
}

// NewAnswerClient returns a client for the Answer from the given config.
func NewAnswerClient(c config) *AnswerClient {
	return &AnswerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `answer.Hooks(f(g(h())))`.
func (c *AnswerClient) Use(hooks ...Hook) {
	c.hooks.Answer = append(c.hooks.Answer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `answer.Intercept(f(g(h())))`.
func (c *AnswerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Answer = append(c.inters.Answer, interceptors...)
}

// Create returns a builder for creating a Answer entity.
func (c *AnswerClient) Create() *AnswerCreate {
	mutation := newAnswerMutation(c.config, OpCreate)
	return &AnswerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Answer entities.
func (c *AnswerClient) CreateBulk(builders ...*AnswerCreate) *AnswerCreateBulk {
	return &AnswerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AnswerClient) MapCreateBulk(slice any, setFunc func(*AnswerCreate, int)) *AnswerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AnswerCreateBulk{err: fmt.Errorf("calling to AnswerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AnswerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AnswerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Answer.
func (c *AnswerClient) Update() *AnswerUpdate {
	mutation := newAnswerMutation(c.config, OpUpdate)
	return &AnswerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnswerClient) UpdateOne(a *Answer) *AnswerUpdateOne {
	mutation := newAnswerMutation(c.config, OpUpdateOne, withAnswer(a))
	return &AnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnswerClient) UpdateOneID(id int) *AnswerUpdateOne {
	mutation := newAnswerMutation(c.config, OpUpdateOne, withAnswerID(id))
	return &AnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Answer.
func (c *AnswerClient) Delete() *AnswerDelete {
	mutation := newAnswerMutation(c.config, OpDelete)
	return &AnswerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AnswerClient) DeleteOne(a *Answer) *AnswerDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AnswerClient) DeleteOneID(id int) *AnswerDeleteOne {
	builder := c.Delete().Where(answer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnswerDeleteOne{builder}
}

// Query returns a query builder for Answer.
func (c *AnswerClient) Query() *AnswerQuery {
	return &AnswerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAnswer},
		inters: c.Interceptors(),
	}
}

// Get returns a Answer entity by its id.
func (c *AnswerClient) Get(ctx context.Context, id int) (*Answer, error) {
	return c.Query().Where(answer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnswerClient) GetX(ctx context.Context, id int) *Answer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQuestion queries the question edge of a Answer.
func (c *AnswerClient) QueryQuestion(a *Answer) *QuestionQuery {
	query := (&QuestionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(answer.Table, answer.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, answer.QuestionTable, answer.QuestionColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Answer.
func (c *AnswerClient) QueryAuthor(a *Answer) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(answer.Table, answer.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, answer.AuthorTable, answer.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AnswerClient) Hooks() []Hook {
	return c.hooks.Answer
}

// Interceptors returns the client interceptors.
func (c *AnswerClient) Interceptors() []Interceptor {
	return c.inters.Answer
}

func (c *AnswerClient) mutate(ctx context.Context, m *AnswerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AnswerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AnswerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AnswerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Answer mutation op: %q", m.Op())
	}
}

// QuestionClient is a client for the Question schema.
type QuestionClient struct {
	config
}

// NewQuestionClient returns a client for the Question from the given config.
func NewQuestionClient(c config) *QuestionClient {
	return &QuestionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `question.Hooks(f(g(h())))`.
func (c *QuestionClient) Use(hooks ...Hook) {
	c.hooks.Question = append(c.hooks.Question, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `question.Intercept(f(g(h())))`.
func (c *QuestionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Question = append(c.inters.Question, interceptors...)
}

// Create returns a builder for creating a Question entity.
func (c *QuestionClient) Create() *QuestionCreate {
	mutation := newQuestionMutation(c.config, OpCreate)
	return &QuestionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Question entities.
func (c *QuestionClient) CreateBulk(builders ...*QuestionCreate) *QuestionCreateBulk {
	return &QuestionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *QuestionClient) MapCreateBulk(slice any, setFunc func(*QuestionCreate, int)) *QuestionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &QuestionCreateBulk{err: fmt.Errorf("calling to QuestionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*QuestionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &QuestionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Question.
func (c *QuestionClient) Update() *QuestionUpdate {
	mutation := newQuestionMutation(c.config, OpUpdate)
	return &QuestionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *QuestionClient) UpdateOne(q *Question) *QuestionUpdateOne {
	mutation := newQuestionMutation(c.config, OpUpdateOne, withQuestion(q))
	return &QuestionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *QuestionClient) UpdateOneID(id int) *QuestionUpdateOne {
	mutation := newQuestionMutation(c.config, OpUpdateOne, withQuestionID(id))
	return &QuestionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Question.
func (c *QuestionClient) Delete() *QuestionDelete {
	mutation := newQuestionMutation(c.config, OpDelete)
	return &QuestionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *QuestionClient) DeleteOne(q *Question) *QuestionDeleteOne {
	return c.DeleteOneID(q.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *QuestionClient) DeleteOneID(id int) *QuestionDeleteOne {
	builder := c.Delete().Where(question.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &QuestionDeleteOne{builder}
}

// Query returns a query builder for Question.
func (c *QuestionClient) Query() *QuestionQuery {
	return &QuestionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeQuestion},
		inters: c.Interceptors(),
	}
}

// Get returns a Question entity by its id.
func (c *QuestionClient) Get(ctx context.Context, id int) (*Question, error) {
	return c.Query().Where(question.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *QuestionClient) GetX(ctx context.Context, id int) *Question {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAnswers queries the answers edge of a Question.
func (c *QuestionClient) QueryAnswers(q *Question) *AnswerQuery {
	query := (&AnswerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, question.AnswersTable, question.AnswersColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Question.
func (c *QuestionClient) QueryAuthor(q *Question) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, question.AuthorTable, question.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a Question.
func (c *QuestionClient) QueryTags(q *Question) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, question.TagsTable, question.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryQuestionTag queries the question_tag edge of a Question.
func (c *QuestionClient) QueryQuestionTag(q *Question) *QuestionTagQuery {
	query := (&QuestionTagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(questiontag.Table, questiontag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, question.QuestionTagTable, question.QuestionTagColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *QuestionClient) Hooks() []Hook {
	return c.hooks.Question
}

// Interceptors returns the client interceptors.
func (c *QuestionClient) Interceptors() []Interceptor {
	return c.inters.Question
}

func (c *QuestionClient) mutate(ctx context.Context, m *QuestionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&QuestionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&QuestionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&QuestionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&QuestionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Question mutation op: %q", m.Op())
	}
}

// QuestionTagClient is a client for the QuestionTag schema.
type QuestionTagClient struct {
	config
}

// NewQuestionTagClient returns a client for the QuestionTag from the given config.
func NewQuestionTagClient(c config) *QuestionTagClient {
	return &QuestionTagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `questiontag.Hooks(f(g(h())))`.
func (c *QuestionTagClient) Use(hooks ...Hook) {
	c.hooks.QuestionTag = append(c.hooks.QuestionTag, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `questiontag.Intercept(f(g(h())))`.
func (c *QuestionTagClient) Intercept(interceptors ...Interceptor) {
	c.inters.QuestionTag = append(c.inters.QuestionTag, interceptors...)
}

// Create returns a builder for creating a QuestionTag entity.
func (c *QuestionTagClient) Create() *QuestionTagCreate {
	mutation := newQuestionTagMutation(c.config, OpCreate)
	return &QuestionTagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of QuestionTag entities.
func (c *QuestionTagClient) CreateBulk(builders ...*QuestionTagCreate) *QuestionTagCreateBulk {
	return &QuestionTagCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *QuestionTagClient) MapCreateBulk(slice any, setFunc func(*QuestionTagCreate, int)) *QuestionTagCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &QuestionTagCreateBulk{err: fmt.Errorf("calling to QuestionTagClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*QuestionTagCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &QuestionTagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for QuestionTag.
func (c *QuestionTagClient) Update() *QuestionTagUpdate {
	mutation := newQuestionTagMutation(c.config, OpUpdate)
	return &QuestionTagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *QuestionTagClient) UpdateOne(qt *QuestionTag) *QuestionTagUpdateOne {
	mutation := newQuestionTagMutation(c.config, OpUpdateOne, withQuestionTag(qt))
	return &QuestionTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *QuestionTagClient) UpdateOneID(id int) *QuestionTagUpdateOne {
	mutation := newQuestionTagMutation(c.config, OpUpdateOne, withQuestionTagID(id))
	return &QuestionTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for QuestionTag.
func (c *QuestionTagClient) Delete() *QuestionTagDelete {
	mutation := newQuestionTagMutation(c.config, OpDelete)
	return &QuestionTagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *QuestionTagClient) DeleteOne(qt *QuestionTag) *QuestionTagDeleteOne {
	return c.DeleteOneID(qt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *QuestionTagClient) DeleteOneID(id int) *QuestionTagDeleteOne {
	builder := c.Delete().Where(questiontag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &QuestionTagDeleteOne{builder}
}

// Query returns a query builder for QuestionTag.
func (c *QuestionTagClient) Query() *QuestionTagQuery {
	return &QuestionTagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeQuestionTag},
		inters: c.Interceptors(),
	}
}

// Get returns a QuestionTag entity by its id.
func (c *QuestionTagClient) Get(ctx context.Context, id int) (*QuestionTag, error) {
	return c.Query().Where(questiontag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *QuestionTagClient) GetX(ctx context.Context, id int) *QuestionTag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQuestion queries the question edge of a QuestionTag.
func (c *QuestionTagClient) QueryQuestion(qt *QuestionTag) *QuestionQuery {
	query := (&QuestionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := qt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(questiontag.Table, questiontag.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, questiontag.QuestionTable, questiontag.QuestionColumn),
		)
		fromV = sqlgraph.Neighbors(qt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTag queries the tag edge of a QuestionTag.
func (c *QuestionTagClient) QueryTag(qt *QuestionTag) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := qt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(questiontag.Table, questiontag.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, questiontag.TagTable, questiontag.TagColumn),
		)
		fromV = sqlgraph.Neighbors(qt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *QuestionTagClient) Hooks() []Hook {
	return c.hooks.QuestionTag
}

// Interceptors returns the client interceptors.
func (c *QuestionTagClient) Interceptors() []Interceptor {
	return c.inters.QuestionTag
}

func (c *QuestionTagClient) mutate(ctx context.Context, m *QuestionTagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&QuestionTagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&QuestionTagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&QuestionTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&QuestionTagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown QuestionTag mutation op: %q", m.Op())
	}
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tag.Intercept(f(g(h())))`.
func (c *TagClient) Intercept(interceptors ...Interceptor) {
	c.inters.Tag = append(c.inters.Tag, interceptors...)
}

// Create returns a builder for creating a Tag entity.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TagClient) MapCreateBulk(slice any, setFunc func(*TagCreate, int)) *TagCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TagCreateBulk{err: fmt.Errorf("calling to TagClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TagCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTag},
		inters: c.Interceptors(),
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQuestions queries the questions edge of a Tag.
func (c *TagClient) QueryQuestions(t *Tag) *QuestionQuery {
	query := (&QuestionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.QuestionsTable, tag.QuestionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTagQuestion queries the tag_question edge of a Tag.
func (c *TagClient) QueryTagQuestion(t *Tag) *QuestionTagQuery {
	query := (&QuestionTagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(questiontag.Table, questiontag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, tag.TagQuestionTable, tag.TagQuestionColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}

// Interceptors returns the client interceptors.
func (c *TagClient) Interceptors() []Interceptor {
	return c.inters.Tag
}

func (c *TagClient) mutate(ctx context.Context, m *TagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Tag mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQuestions queries the questions edge of a User.
func (c *UserClient) QueryQuestions(u *User) *QuestionQuery {
	query := (&QuestionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.QuestionsTable, user.QuestionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAnswers queries the answers edge of a User.
func (c *UserClient) QueryAnswers(u *User) *AnswerQuery {
	query := (&AnswerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.AnswersTable, user.AnswersColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a User.
func (c *UserClient) QueryTags(u *User) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TagsTable, user.TagsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Answer, Question, QuestionTag, Tag, User []ent.Hook
	}
	inters struct {
		Answer, Question, QuestionTag, Tag, User []ent.Interceptor
	}
)
