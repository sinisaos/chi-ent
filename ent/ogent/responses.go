// Code generated by ent, DO NOT EDIT.

package ogent

import "github.com/sinisaos/chi-ent/ent"

func NewAnswerCreate(e *ent.Answer) *AnswerCreate {
	if e == nil {
		return nil
	}
	var ret AnswerCreate
	ret.ID = e.ID
	ret.Content = e.Content
	ret.Likes = e.Likes
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.IsAcceptedAnswer = e.IsAcceptedAnswer
	return &ret
}

func NewAnswerCreates(es []*ent.Answer) []AnswerCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]AnswerCreate, len(es))
	for i, e := range es {
		r[i] = NewAnswerCreate(e).Elem()
	}
	return r
}

func (a *AnswerCreate) Elem() AnswerCreate {
	if a == nil {
		return AnswerCreate{}
	}
	return *a
}

func NewAnswerList(e *ent.Answer) *AnswerList {
	if e == nil {
		return nil
	}
	var ret AnswerList
	ret.ID = e.ID
	ret.Content = e.Content
	ret.Likes = e.Likes
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.IsAcceptedAnswer = e.IsAcceptedAnswer
	return &ret
}

func NewAnswerLists(es []*ent.Answer) []AnswerList {
	if len(es) == 0 {
		return nil
	}
	r := make([]AnswerList, len(es))
	for i, e := range es {
		r[i] = NewAnswerList(e).Elem()
	}
	return r
}

func (a *AnswerList) Elem() AnswerList {
	if a == nil {
		return AnswerList{}
	}
	return *a
}

func NewAnswerRead(e *ent.Answer) *AnswerRead {
	if e == nil {
		return nil
	}
	var ret AnswerRead
	ret.ID = e.ID
	ret.Content = e.Content
	ret.Likes = e.Likes
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.IsAcceptedAnswer = e.IsAcceptedAnswer
	return &ret
}

func NewAnswerReads(es []*ent.Answer) []AnswerRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]AnswerRead, len(es))
	for i, e := range es {
		r[i] = NewAnswerRead(e).Elem()
	}
	return r
}

func (a *AnswerRead) Elem() AnswerRead {
	if a == nil {
		return AnswerRead{}
	}
	return *a
}

func NewAnswerUpdate(e *ent.Answer) *AnswerUpdate {
	if e == nil {
		return nil
	}
	var ret AnswerUpdate
	ret.ID = e.ID
	ret.Content = e.Content
	ret.Likes = e.Likes
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.IsAcceptedAnswer = e.IsAcceptedAnswer
	return &ret
}

func NewAnswerUpdates(es []*ent.Answer) []AnswerUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]AnswerUpdate, len(es))
	for i, e := range es {
		r[i] = NewAnswerUpdate(e).Elem()
	}
	return r
}

func (a *AnswerUpdate) Elem() AnswerUpdate {
	if a == nil {
		return AnswerUpdate{}
	}
	return *a
}

func NewAnswerAuthorRead(e *ent.User) *AnswerAuthorRead {
	if e == nil {
		return nil
	}
	var ret AnswerAuthorRead
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	ret.LastLogin = e.LastLogin
	return &ret
}

func NewAnswerAuthorReads(es []*ent.User) []AnswerAuthorRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]AnswerAuthorRead, len(es))
	for i, e := range es {
		r[i] = NewAnswerAuthorRead(e).Elem()
	}
	return r
}

func (u *AnswerAuthorRead) Elem() AnswerAuthorRead {
	if u == nil {
		return AnswerAuthorRead{}
	}
	return *u
}

func NewAnswerQuestionRead(e *ent.Question) *AnswerQuestionRead {
	if e == nil {
		return nil
	}
	var ret AnswerQuestionRead
	ret.ID = e.ID
	ret.Title = e.Title
	ret.Slug = e.Slug
	ret.Content = e.Content
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.Views = e.Views
	ret.Likes = e.Likes
	return &ret
}

func NewAnswerQuestionReads(es []*ent.Question) []AnswerQuestionRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]AnswerQuestionRead, len(es))
	for i, e := range es {
		r[i] = NewAnswerQuestionRead(e).Elem()
	}
	return r
}

func (q *AnswerQuestionRead) Elem() AnswerQuestionRead {
	if q == nil {
		return AnswerQuestionRead{}
	}
	return *q
}

func NewQuestionCreate(e *ent.Question) *QuestionCreate {
	if e == nil {
		return nil
	}
	var ret QuestionCreate
	ret.ID = e.ID
	ret.Title = e.Title
	ret.Slug = e.Slug
	ret.Content = e.Content
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.Views = e.Views
	ret.Likes = e.Likes
	return &ret
}

func NewQuestionCreates(es []*ent.Question) []QuestionCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]QuestionCreate, len(es))
	for i, e := range es {
		r[i] = NewQuestionCreate(e).Elem()
	}
	return r
}

func (q *QuestionCreate) Elem() QuestionCreate {
	if q == nil {
		return QuestionCreate{}
	}
	return *q
}

func NewQuestionList(e *ent.Question) *QuestionList {
	if e == nil {
		return nil
	}
	var ret QuestionList
	ret.ID = e.ID
	ret.Title = e.Title
	ret.Slug = e.Slug
	ret.Content = e.Content
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.Views = e.Views
	ret.Likes = e.Likes
	return &ret
}

func NewQuestionLists(es []*ent.Question) []QuestionList {
	if len(es) == 0 {
		return nil
	}
	r := make([]QuestionList, len(es))
	for i, e := range es {
		r[i] = NewQuestionList(e).Elem()
	}
	return r
}

func (q *QuestionList) Elem() QuestionList {
	if q == nil {
		return QuestionList{}
	}
	return *q
}

func NewQuestionRead(e *ent.Question) *QuestionRead {
	if e == nil {
		return nil
	}
	var ret QuestionRead
	ret.ID = e.ID
	ret.Title = e.Title
	ret.Slug = e.Slug
	ret.Content = e.Content
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.Views = e.Views
	ret.Likes = e.Likes
	return &ret
}

func NewQuestionReads(es []*ent.Question) []QuestionRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]QuestionRead, len(es))
	for i, e := range es {
		r[i] = NewQuestionRead(e).Elem()
	}
	return r
}

func (q *QuestionRead) Elem() QuestionRead {
	if q == nil {
		return QuestionRead{}
	}
	return *q
}

func NewQuestionUpdate(e *ent.Question) *QuestionUpdate {
	if e == nil {
		return nil
	}
	var ret QuestionUpdate
	ret.ID = e.ID
	ret.Title = e.Title
	ret.Slug = e.Slug
	ret.Content = e.Content
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.Views = e.Views
	ret.Likes = e.Likes
	return &ret
}

func NewQuestionUpdates(es []*ent.Question) []QuestionUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]QuestionUpdate, len(es))
	for i, e := range es {
		r[i] = NewQuestionUpdate(e).Elem()
	}
	return r
}

func (q *QuestionUpdate) Elem() QuestionUpdate {
	if q == nil {
		return QuestionUpdate{}
	}
	return *q
}

func NewQuestionAnswersList(e *ent.Answer) *QuestionAnswersList {
	if e == nil {
		return nil
	}
	var ret QuestionAnswersList
	ret.ID = e.ID
	ret.Content = e.Content
	ret.Likes = e.Likes
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.IsAcceptedAnswer = e.IsAcceptedAnswer
	return &ret
}

func NewQuestionAnswersLists(es []*ent.Answer) []QuestionAnswersList {
	if len(es) == 0 {
		return nil
	}
	r := make([]QuestionAnswersList, len(es))
	for i, e := range es {
		r[i] = NewQuestionAnswersList(e).Elem()
	}
	return r
}

func (a *QuestionAnswersList) Elem() QuestionAnswersList {
	if a == nil {
		return QuestionAnswersList{}
	}
	return *a
}

func NewQuestionAuthorRead(e *ent.User) *QuestionAuthorRead {
	if e == nil {
		return nil
	}
	var ret QuestionAuthorRead
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	ret.LastLogin = e.LastLogin
	return &ret
}

func NewQuestionAuthorReads(es []*ent.User) []QuestionAuthorRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]QuestionAuthorRead, len(es))
	for i, e := range es {
		r[i] = NewQuestionAuthorRead(e).Elem()
	}
	return r
}

func (u *QuestionAuthorRead) Elem() QuestionAuthorRead {
	if u == nil {
		return QuestionAuthorRead{}
	}
	return *u
}

func NewQuestionTagsList(e *ent.Tag) *QuestionTagsList {
	if e == nil {
		return nil
	}
	var ret QuestionTagsList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewQuestionTagsLists(es []*ent.Tag) []QuestionTagsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]QuestionTagsList, len(es))
	for i, e := range es {
		r[i] = NewQuestionTagsList(e).Elem()
	}
	return r
}

func (t *QuestionTagsList) Elem() QuestionTagsList {
	if t == nil {
		return QuestionTagsList{}
	}
	return *t
}

func NewTagCreate(e *ent.Tag) *TagCreate {
	if e == nil {
		return nil
	}
	var ret TagCreate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewTagCreates(es []*ent.Tag) []TagCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]TagCreate, len(es))
	for i, e := range es {
		r[i] = NewTagCreate(e).Elem()
	}
	return r
}

func (t *TagCreate) Elem() TagCreate {
	if t == nil {
		return TagCreate{}
	}
	return *t
}

func NewTagList(e *ent.Tag) *TagList {
	if e == nil {
		return nil
	}
	var ret TagList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewTagLists(es []*ent.Tag) []TagList {
	if len(es) == 0 {
		return nil
	}
	r := make([]TagList, len(es))
	for i, e := range es {
		r[i] = NewTagList(e).Elem()
	}
	return r
}

func (t *TagList) Elem() TagList {
	if t == nil {
		return TagList{}
	}
	return *t
}

func NewTagRead(e *ent.Tag) *TagRead {
	if e == nil {
		return nil
	}
	var ret TagRead
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewTagReads(es []*ent.Tag) []TagRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]TagRead, len(es))
	for i, e := range es {
		r[i] = NewTagRead(e).Elem()
	}
	return r
}

func (t *TagRead) Elem() TagRead {
	if t == nil {
		return TagRead{}
	}
	return *t
}

func NewTagUpdate(e *ent.Tag) *TagUpdate {
	if e == nil {
		return nil
	}
	var ret TagUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewTagUpdates(es []*ent.Tag) []TagUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]TagUpdate, len(es))
	for i, e := range es {
		r[i] = NewTagUpdate(e).Elem()
	}
	return r
}

func (t *TagUpdate) Elem() TagUpdate {
	if t == nil {
		return TagUpdate{}
	}
	return *t
}

func NewTagQuestionsList(e *ent.Question) *TagQuestionsList {
	if e == nil {
		return nil
	}
	var ret TagQuestionsList
	ret.ID = e.ID
	ret.Title = e.Title
	ret.Slug = e.Slug
	ret.Content = e.Content
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.Views = e.Views
	ret.Likes = e.Likes
	return &ret
}

func NewTagQuestionsLists(es []*ent.Question) []TagQuestionsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]TagQuestionsList, len(es))
	for i, e := range es {
		r[i] = NewTagQuestionsList(e).Elem()
	}
	return r
}

func (q *TagQuestionsList) Elem() TagQuestionsList {
	if q == nil {
		return TagQuestionsList{}
	}
	return *q
}

func NewUserCreate(e *ent.User) *UserCreate {
	if e == nil {
		return nil
	}
	var ret UserCreate
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	ret.LastLogin = e.LastLogin
	return &ret
}

func NewUserCreates(es []*ent.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u == nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *ent.User) *UserList {
	if e == nil {
		return nil
	}
	var ret UserList
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	ret.LastLogin = e.LastLogin
	return &ret
}

func NewUserLists(es []*ent.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u == nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *ent.User) *UserRead {
	if e == nil {
		return nil
	}
	var ret UserRead
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	ret.LastLogin = e.LastLogin
	return &ret
}

func NewUserReads(es []*ent.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u == nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *ent.User) *UserUpdate {
	if e == nil {
		return nil
	}
	var ret UserUpdate
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	ret.LastLogin = e.LastLogin
	return &ret
}

func NewUserUpdates(es []*ent.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u == nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserAnswersList(e *ent.Answer) *UserAnswersList {
	if e == nil {
		return nil
	}
	var ret UserAnswersList
	ret.ID = e.ID
	ret.Content = e.Content
	ret.Likes = e.Likes
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.IsAcceptedAnswer = e.IsAcceptedAnswer
	return &ret
}

func NewUserAnswersLists(es []*ent.Answer) []UserAnswersList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserAnswersList, len(es))
	for i, e := range es {
		r[i] = NewUserAnswersList(e).Elem()
	}
	return r
}

func (a *UserAnswersList) Elem() UserAnswersList {
	if a == nil {
		return UserAnswersList{}
	}
	return *a
}

func NewUserQuestionsList(e *ent.Question) *UserQuestionsList {
	if e == nil {
		return nil
	}
	var ret UserQuestionsList
	ret.ID = e.ID
	ret.Title = e.Title
	ret.Slug = e.Slug
	ret.Content = e.Content
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	ret.Views = e.Views
	ret.Likes = e.Likes
	return &ret
}

func NewUserQuestionsLists(es []*ent.Question) []UserQuestionsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserQuestionsList, len(es))
	for i, e := range es {
		r[i] = NewUserQuestionsList(e).Elem()
	}
	return r
}

func (q *UserQuestionsList) Elem() UserQuestionsList {
	if q == nil {
		return UserQuestionsList{}
	}
	return *q
}

func NewUserTagsList(e *ent.Tag) *UserTagsList {
	if e == nil {
		return nil
	}
	var ret UserTagsList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewUserTagsLists(es []*ent.Tag) []UserTagsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserTagsList, len(es))
	for i, e := range es {
		r[i] = NewUserTagsList(e).Elem()
	}
	return r
}

func (t *UserTagsList) Elem() UserTagsList {
	if t == nil {
		return UserTagsList{}
	}
	return *t
}
