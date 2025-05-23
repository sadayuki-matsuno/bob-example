// Code generated by sadayuki-matsuno bob-ex. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"slices"
	"time"

	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
	"github.com/stephenafamo/bob/expr"
	"github.com/stephenafamo/bob/mods"
	"github.com/stephenafamo/bob/orm"
	"github.com/stephenafamo/bob/types/pgtypes"
)

// Comment is an object representing the database table.
type Comment struct {
	ID        int32               `db:"id,pk" `
	PostID    int32               `db:"post_id" `
	UserID    int32               `db:"user_id" `
	Body      string              `db:"body" `
	CreatedAt null.Val[time.Time] `db:"created_at" `
	UpdatedAt null.Val[time.Time] `db:"updated_at" `

	R commentR `db:"-" `
}

// CommentSlice is an alias for a slice of pointers to Comment.
// This should almost always be used instead of []*Comment.
type CommentSlice []*Comment

// Comments contains methods to work with the comments table
var Comments = psql.NewTablex[*Comment, CommentSlice, *CommentSetter]("", "comments")

// CommentsQuery is a query on the comments table
type CommentsQuery = *psql.ViewQuery[*Comment, CommentSlice]

// commentR is where relationships are stored.
type commentR struct {
	Post *Post // comments.fk_comments_post
	User *User // comments.fk_comments_user
}

type commentColumnNames struct {
	ID        string
	PostID    string
	UserID    string
	Body      string
	CreatedAt string
	UpdatedAt string
}

var CommentColumns = buildCommentColumns("comments")

type commentColumns struct {
	tableAlias string
	ID         psql.Expression
	PostID     psql.Expression
	UserID     psql.Expression
	Body       psql.Expression
	CreatedAt  psql.Expression
	UpdatedAt  psql.Expression
}

func (c commentColumns) Alias() string {
	return c.tableAlias
}

func (commentColumns) AliasedAs(alias string) commentColumns {
	return buildCommentColumns(alias)
}

func buildCommentColumns(alias string) commentColumns {
	return commentColumns{
		tableAlias: alias,
		ID:         psql.Quote(alias, "id"),
		PostID:     psql.Quote(alias, "post_id"),
		UserID:     psql.Quote(alias, "user_id"),
		Body:       psql.Quote(alias, "body"),
		CreatedAt:  psql.Quote(alias, "created_at"),
		UpdatedAt:  psql.Quote(alias, "updated_at"),
	}
}

type commentWhere[Q psql.Filterable] struct {
	ID        psql.WhereMod[Q, int32]
	PostID    psql.WhereMod[Q, int32]
	UserID    psql.WhereMod[Q, int32]
	Body      psql.WhereMod[Q, string]
	CreatedAt psql.WhereNullMod[Q, time.Time]
	UpdatedAt psql.WhereNullMod[Q, time.Time]
}

func (commentWhere[Q]) AliasedAs(alias string) commentWhere[Q] {
	return buildCommentWhere[Q](buildCommentColumns(alias))
}

func buildCommentWhere[Q psql.Filterable](cols commentColumns) commentWhere[Q] {
	return commentWhere[Q]{
		ID:        psql.Where[Q, int32](cols.ID),
		PostID:    psql.Where[Q, int32](cols.PostID),
		UserID:    psql.Where[Q, int32](cols.UserID),
		Body:      psql.Where[Q, string](cols.Body),
		CreatedAt: psql.WhereNull[Q, time.Time](cols.CreatedAt),
		UpdatedAt: psql.WhereNull[Q, time.Time](cols.UpdatedAt),
	}
}

var CommentErrors = &commentErrors{
	ErrUniqueCommentsPkey: &UniqueConstraintError{s: "comments_pkey"},
}

type commentErrors struct {
	ErrUniqueCommentsPkey *UniqueConstraintError
}

// CommentSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type CommentSetter struct {
	ID        omit.Val[int32]         `db:"id,pk" `
	PostID    omit.Val[int32]         `db:"post_id" `
	UserID    omit.Val[int32]         `db:"user_id" `
	Body      omit.Val[string]        `db:"body" `
	CreatedAt omitnull.Val[time.Time] `db:"created_at" `
	UpdatedAt omitnull.Val[time.Time] `db:"updated_at" `
}

func (s CommentSetter) SetColumns() []string {
	vals := make([]string, 0, 6)
	if !s.ID.IsUnset() {
		vals = append(vals, "id")
	}

	if !s.PostID.IsUnset() {
		vals = append(vals, "post_id")
	}

	if !s.UserID.IsUnset() {
		vals = append(vals, "user_id")
	}

	if !s.Body.IsUnset() {
		vals = append(vals, "body")
	}

	if !s.CreatedAt.IsUnset() {
		vals = append(vals, "created_at")
	}

	if !s.UpdatedAt.IsUnset() {
		vals = append(vals, "updated_at")
	}

	return vals
}

func (s CommentSetter) Overwrite(t *Comment) {
	if !s.ID.IsUnset() {
		t.ID, _ = s.ID.Get()
	}
	if !s.PostID.IsUnset() {
		t.PostID, _ = s.PostID.Get()
	}
	if !s.UserID.IsUnset() {
		t.UserID, _ = s.UserID.Get()
	}
	if !s.Body.IsUnset() {
		t.Body, _ = s.Body.Get()
	}
	if !s.CreatedAt.IsUnset() {
		t.CreatedAt, _ = s.CreatedAt.GetNull()
	}
	if !s.UpdatedAt.IsUnset() {
		t.UpdatedAt, _ = s.UpdatedAt.GetNull()
	}
}

func (s *CommentSetter) Apply(q *dialect.InsertQuery) {
	q.AppendHooks(func(ctx context.Context, exec bob.Executor) (context.Context, error) {
		return Comments.BeforeInsertHooks.RunHooks(ctx, exec, s)
	})

	q.AppendValues(bob.ExpressionFunc(func(ctx context.Context, w io.Writer, d bob.Dialect, start int) ([]any, error) {
		vals := make([]bob.Expression, 6)
		if s.ID.IsUnset() {
			vals[0] = psql.Raw("DEFAULT")
		} else {
			vals[0] = psql.Arg(s.ID)
		}

		if s.PostID.IsUnset() {
			vals[1] = psql.Raw("DEFAULT")
		} else {
			vals[1] = psql.Arg(s.PostID)
		}

		if s.UserID.IsUnset() {
			vals[2] = psql.Raw("DEFAULT")
		} else {
			vals[2] = psql.Arg(s.UserID)
		}

		if s.Body.IsUnset() {
			vals[3] = psql.Raw("DEFAULT")
		} else {
			vals[3] = psql.Arg(s.Body)
		}

		if s.CreatedAt.IsUnset() {
			vals[4] = psql.Raw("DEFAULT")
		} else {
			vals[4] = psql.Arg(s.CreatedAt)
		}

		if s.UpdatedAt.IsUnset() {
			vals[5] = psql.Raw("DEFAULT")
		} else {
			vals[5] = psql.Arg(s.UpdatedAt)
		}

		return bob.ExpressSlice(ctx, w, d, start, vals, "", ", ", "")
	}))
}

func (s CommentSetter) UpdateMod() bob.Mod[*dialect.UpdateQuery] {
	return um.Set(s.Expressions()...)
}

func (s CommentSetter) Expressions(prefix ...string) []bob.Expression {
	exprs := make([]bob.Expression, 0, 6)

	if !s.ID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "id")...),
			psql.Arg(s.ID),
		}})
	}

	if !s.PostID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "post_id")...),
			psql.Arg(s.PostID),
		}})
	}

	if !s.UserID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "user_id")...),
			psql.Arg(s.UserID),
		}})
	}

	if !s.Body.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "body")...),
			psql.Arg(s.Body),
		}})
	}

	if !s.CreatedAt.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "created_at")...),
			psql.Arg(s.CreatedAt),
		}})
	}

	if !s.UpdatedAt.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "updated_at")...),
			psql.Arg(s.UpdatedAt),
		}})
	}

	return exprs
}

// FindComment retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindComment(ctx context.Context, exec bob.Executor, IDPK int32, cols ...string) (*Comment, error) {
	if len(cols) == 0 {
		return Comments.Query(
			SelectWhere.Comments.ID.EQ(IDPK),
		).One(ctx, exec)
	}

	return Comments.Query(
		SelectWhere.Comments.ID.EQ(IDPK),
		sm.Columns(Comments.Columns().Only(cols...)),
	).One(ctx, exec)
}

// CommentExists checks the presence of a single record by primary key
func CommentExists(ctx context.Context, exec bob.Executor, IDPK int32) (bool, error) {
	return Comments.Query(
		SelectWhere.Comments.ID.EQ(IDPK),
	).Exists(ctx, exec)
}

// AfterQueryHook is called after Comment is retrieved from the database
func (o *Comment) AfterQueryHook(ctx context.Context, exec bob.Executor, queryType bob.QueryType) error {
	var err error

	switch queryType {
	case bob.QueryTypeSelect:
		ctx, err = Comments.AfterSelectHooks.RunHooks(ctx, exec, CommentSlice{o})
	case bob.QueryTypeInsert:
		ctx, err = Comments.AfterInsertHooks.RunHooks(ctx, exec, CommentSlice{o})
	case bob.QueryTypeUpdate:
		ctx, err = Comments.AfterUpdateHooks.RunHooks(ctx, exec, CommentSlice{o})
	case bob.QueryTypeDelete:
		ctx, err = Comments.AfterDeleteHooks.RunHooks(ctx, exec, CommentSlice{o})
	}

	return err
}

// PrimaryKeyVals returns the primary key values of the Comment
func (o *Comment) PrimaryKeyVals() bob.Expression {
	return psql.Arg(o.ID)
}

func (o *Comment) pkEQ() dialect.Expression {
	return psql.Quote("comments", "id").EQ(bob.ExpressionFunc(func(ctx context.Context, w io.Writer, d bob.Dialect, start int) ([]any, error) {
		return o.PrimaryKeyVals().WriteSQL(ctx, w, d, start)
	}))
}

// Update uses an executor to update the Comment
func (o *Comment) Update(ctx context.Context, exec bob.Executor, s *CommentSetter) error {
	v, err := Comments.Update(s.UpdateMod(), um.Where(o.pkEQ())).One(ctx, exec)
	if err != nil {
		return err
	}

	o.R = v.R
	*o = *v

	return nil
}

// Delete deletes a single Comment record with an executor
func (o *Comment) Delete(ctx context.Context, exec bob.Executor) error {
	_, err := Comments.Delete(dm.Where(o.pkEQ())).Exec(ctx, exec)
	return err
}

// Reload refreshes the Comment using the executor
func (o *Comment) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := Comments.Query(
		SelectWhere.Comments.ID.EQ(o.ID),
	).One(ctx, exec)
	if err != nil {
		return err
	}
	o2.R = o.R
	*o = *o2

	return nil
}

// AfterQueryHook is called after CommentSlice is retrieved from the database
func (o CommentSlice) AfterQueryHook(ctx context.Context, exec bob.Executor, queryType bob.QueryType) error {
	var err error

	switch queryType {
	case bob.QueryTypeSelect:
		ctx, err = Comments.AfterSelectHooks.RunHooks(ctx, exec, o)
	case bob.QueryTypeInsert:
		ctx, err = Comments.AfterInsertHooks.RunHooks(ctx, exec, o)
	case bob.QueryTypeUpdate:
		ctx, err = Comments.AfterUpdateHooks.RunHooks(ctx, exec, o)
	case bob.QueryTypeDelete:
		ctx, err = Comments.AfterDeleteHooks.RunHooks(ctx, exec, o)
	}

	return err
}

func (o CommentSlice) pkIN() dialect.Expression {
	if len(o) == 0 {
		return psql.Raw("NULL")
	}

	return psql.Quote("comments", "id").In(bob.ExpressionFunc(func(ctx context.Context, w io.Writer, d bob.Dialect, start int) ([]any, error) {
		pkPairs := make([]bob.Expression, len(o))
		for i, row := range o {
			pkPairs[i] = row.PrimaryKeyVals()
		}
		return bob.ExpressSlice(ctx, w, d, start, pkPairs, "", ", ", "")
	}))
}

// copyMatchingRows finds models in the given slice that have the same primary key
// then it first copies the existing relationships from the old model to the new model
// and then replaces the old model in the slice with the new model
func (o CommentSlice) copyMatchingRows(from ...*Comment) {
	for i, old := range o {
		for _, new := range from {
			if new.ID != old.ID {
				continue
			}
			new.R = old.R
			o[i] = new
			break
		}
	}
}

// UpdateMod modifies an update query with "WHERE primary_key IN (o...)"
func (o CommentSlice) UpdateMod() bob.Mod[*dialect.UpdateQuery] {
	return bob.ModFunc[*dialect.UpdateQuery](func(q *dialect.UpdateQuery) {
		q.AppendHooks(func(ctx context.Context, exec bob.Executor) (context.Context, error) {
			return Comments.BeforeUpdateHooks.RunHooks(ctx, exec, o)
		})

		q.AppendLoader(bob.LoaderFunc(func(ctx context.Context, exec bob.Executor, retrieved any) error {
			var err error
			switch retrieved := retrieved.(type) {
			case *Comment:
				o.copyMatchingRows(retrieved)
			case []*Comment:
				o.copyMatchingRows(retrieved...)
			case CommentSlice:
				o.copyMatchingRows(retrieved...)
			default:
				// If the retrieved value is not a Comment or a slice of Comment
				// then run the AfterUpdateHooks on the slice
				_, err = Comments.AfterUpdateHooks.RunHooks(ctx, exec, o)
			}

			return err
		}))

		q.AppendWhere(o.pkIN())
	})
}

// DeleteMod modifies an delete query with "WHERE primary_key IN (o...)"
func (o CommentSlice) DeleteMod() bob.Mod[*dialect.DeleteQuery] {
	return bob.ModFunc[*dialect.DeleteQuery](func(q *dialect.DeleteQuery) {
		q.AppendHooks(func(ctx context.Context, exec bob.Executor) (context.Context, error) {
			return Comments.BeforeDeleteHooks.RunHooks(ctx, exec, o)
		})

		q.AppendLoader(bob.LoaderFunc(func(ctx context.Context, exec bob.Executor, retrieved any) error {
			var err error
			switch retrieved := retrieved.(type) {
			case *Comment:
				o.copyMatchingRows(retrieved)
			case []*Comment:
				o.copyMatchingRows(retrieved...)
			case CommentSlice:
				o.copyMatchingRows(retrieved...)
			default:
				// If the retrieved value is not a Comment or a slice of Comment
				// then run the AfterDeleteHooks on the slice
				_, err = Comments.AfterDeleteHooks.RunHooks(ctx, exec, o)
			}

			return err
		}))

		q.AppendWhere(o.pkIN())
	})
}

func (o CommentSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals CommentSetter) error {
	if len(o) == 0 {
		return nil
	}

	_, err := Comments.Update(vals.UpdateMod(), o.UpdateMod()).All(ctx, exec)
	return err
}

func (o CommentSlice) DeleteAll(ctx context.Context, exec bob.Executor) error {
	if len(o) == 0 {
		return nil
	}

	_, err := Comments.Delete(o.DeleteMod()).Exec(ctx, exec)
	return err
}

func (o CommentSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	if len(o) == 0 {
		return nil
	}

	o2, err := Comments.Query(sm.Where(o.pkIN())).All(ctx, exec)
	if err != nil {
		return err
	}

	o.copyMatchingRows(o2...)

	return nil
}

type commentJoins[Q dialect.Joinable] struct {
	typ  string
	Post func(context.Context) modAs[Q, postColumns]
	User func(context.Context) modAs[Q, userColumns]
}

func (j commentJoins[Q]) aliasedAs(alias string) commentJoins[Q] {
	return buildCommentJoins[Q](buildCommentColumns(alias), j.typ)
}

func buildCommentJoins[Q dialect.Joinable](cols commentColumns, typ string) commentJoins[Q] {
	return commentJoins[Q]{
		typ:  typ,
		Post: commentsJoinPost[Q](cols, typ),
		User: commentsJoinUser[Q](cols, typ),
	}
}

func commentsJoinPost[Q dialect.Joinable](from commentColumns, typ string) func(context.Context) modAs[Q, postColumns] {
	return func(ctx context.Context) modAs[Q, postColumns] {
		return modAs[Q, postColumns]{
			c: PostColumns,
			f: func(to postColumns) bob.Mod[Q] {
				mods := make(mods.QueryMods[Q], 0, 1)

				{
					mods = append(mods, dialect.Join[Q](typ, Posts.Name().As(to.Alias())).On(
						to.ID.EQ(from.PostID),
					))
				}

				return mods
			},
		}
	}
}

func commentsJoinUser[Q dialect.Joinable](from commentColumns, typ string) func(context.Context) modAs[Q, userColumns] {
	return func(ctx context.Context) modAs[Q, userColumns] {
		return modAs[Q, userColumns]{
			c: UserColumns,
			f: func(to userColumns) bob.Mod[Q] {
				mods := make(mods.QueryMods[Q], 0, 1)

				{
					mods = append(mods, dialect.Join[Q](typ, Users.Name().As(to.Alias())).On(
						to.ID.EQ(from.UserID),
					))
				}

				return mods
			},
		}
	}
}

// Post starts a query for related objects on posts
func (o *Comment) Post(mods ...bob.Mod[*dialect.SelectQuery]) PostsQuery {
	return Posts.Query(append(mods,
		sm.Where(PostColumns.ID.EQ(psql.Arg(o.PostID))),
	)...)
}

func (os CommentSlice) Post(mods ...bob.Mod[*dialect.SelectQuery]) PostsQuery {
	pkPostID := make(pgtypes.Array[int32], len(os))
	for i, o := range os {
		pkPostID[i] = o.PostID
	}
	PKArgExpr := psql.Select(sm.Columns(
		psql.F("unnest", psql.Cast(psql.Arg(pkPostID), "integer[]")),
	))

	return Posts.Query(append(mods,
		sm.Where(psql.Group(PostColumns.ID).In(PKArgExpr)),
	)...)
}

// User starts a query for related objects on users
func (o *Comment) User(mods ...bob.Mod[*dialect.SelectQuery]) UsersQuery {
	return Users.Query(append(mods,
		sm.Where(UserColumns.ID.EQ(psql.Arg(o.UserID))),
	)...)
}

func (os CommentSlice) User(mods ...bob.Mod[*dialect.SelectQuery]) UsersQuery {
	pkUserID := make(pgtypes.Array[int32], len(os))
	for i, o := range os {
		pkUserID[i] = o.UserID
	}
	PKArgExpr := psql.Select(sm.Columns(
		psql.F("unnest", psql.Cast(psql.Arg(pkUserID), "integer[]")),
	))

	return Users.Query(append(mods,
		sm.Where(psql.Group(UserColumns.ID).In(PKArgExpr)),
	)...)
}

func (o *Comment) Preload(name string, retrieved any) error {
	if o == nil {
		return nil
	}

	switch name {
	case "Post":
		rel, ok := retrieved.(*Post)
		if !ok {
			return fmt.Errorf("comment cannot load %T as %q", retrieved, name)
		}

		o.R.Post = rel

		if rel != nil {
			rel.R.Comments = CommentSlice{o}
		}
		return nil
	case "User":
		rel, ok := retrieved.(*User)
		if !ok {
			return fmt.Errorf("comment cannot load %T as %q", retrieved, name)
		}

		o.R.User = rel

		if rel != nil {
			rel.R.Comments = CommentSlice{o}
		}
		return nil
	default:
		return fmt.Errorf("comment has no relationship %q", name)
	}
}

func PreloadCommentPost(opts ...psql.PreloadOption) psql.Preloader {
	return psql.Preload[*Post, PostSlice](orm.Relationship{
		Name: "Post",
		Sides: []orm.RelSide{
			{
				From: TableNames.Comments,
				To:   TableNames.Posts,
				FromColumns: []string{
					ColumnNames.Comments.PostID,
				},
				ToColumns: []string{
					ColumnNames.Posts.ID,
				},
			},
		},
	}, Posts.Columns().Names(), opts...)
}

func ThenLoadCommentPost(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadCommentPost(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load CommentPost", retrieved)
		}

		err := loader.LoadCommentPost(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadCommentPost loads the comment's Post into the .R struct
func (o *Comment) LoadCommentPost(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Post = nil

	related, err := o.Post(mods...).One(ctx, exec)
	if err != nil {
		return err
	}

	related.R.Comments = CommentSlice{o}

	o.R.Post = related
	return nil
}

// LoadCommentPost loads the comment's Post into the .R struct
func (os CommentSlice) LoadCommentPost(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	posts, err := os.Post(mods...).All(ctx, exec)
	if err != nil {
		return err
	}

	for _, o := range os {
		for _, rel := range posts {
			if o.PostID != rel.ID {
				continue
			}

			rel.R.Comments = append(rel.R.Comments, o)

			o.R.Post = rel
			break
		}
	}

	return nil
}

func PreloadCommentUser(opts ...psql.PreloadOption) psql.Preloader {
	return psql.Preload[*User, UserSlice](orm.Relationship{
		Name: "User",
		Sides: []orm.RelSide{
			{
				From: TableNames.Comments,
				To:   TableNames.Users,
				FromColumns: []string{
					ColumnNames.Comments.UserID,
				},
				ToColumns: []string{
					ColumnNames.Users.ID,
				},
			},
		},
	}, Users.Columns().Names(), opts...)
}

func ThenLoadCommentUser(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadCommentUser(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load CommentUser", retrieved)
		}

		err := loader.LoadCommentUser(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadCommentUser loads the comment's User into the .R struct
func (o *Comment) LoadCommentUser(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.User = nil

	related, err := o.User(mods...).One(ctx, exec)
	if err != nil {
		return err
	}

	related.R.Comments = CommentSlice{o}

	o.R.User = related
	return nil
}

// LoadCommentUser loads the comment's User into the .R struct
func (os CommentSlice) LoadCommentUser(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	users, err := os.User(mods...).All(ctx, exec)
	if err != nil {
		return err
	}

	for _, o := range os {
		for _, rel := range users {
			if o.UserID != rel.ID {
				continue
			}

			rel.R.Comments = append(rel.R.Comments, o)

			o.R.User = rel
			break
		}
	}

	return nil
}

func attachCommentPost0(ctx context.Context, exec bob.Executor, count int, comment0 *Comment, post1 *Post) (*Comment, error) {
	setter := &CommentSetter{
		PostID: omit.From(post1.ID),
	}

	err := comment0.Update(ctx, exec, setter)
	if err != nil {
		return nil, fmt.Errorf("attachCommentPost0: %w", err)
	}

	return comment0, nil
}

func (comment0 *Comment) InsertPost(ctx context.Context, exec bob.Executor, related *PostSetter) error {
	post1, err := Posts.Insert(related).One(ctx, exec)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}

	_, err = attachCommentPost0(ctx, exec, 1, comment0, post1)
	if err != nil {
		return err
	}

	comment0.R.Post = post1

	post1.R.Comments = append(post1.R.Comments, comment0)

	return nil
}

func (comment0 *Comment) AttachPost(ctx context.Context, exec bob.Executor, post1 *Post) error {
	var err error

	_, err = attachCommentPost0(ctx, exec, 1, comment0, post1)
	if err != nil {
		return err
	}

	comment0.R.Post = post1

	post1.R.Comments = append(post1.R.Comments, comment0)

	return nil
}

func attachCommentUser0(ctx context.Context, exec bob.Executor, count int, comment0 *Comment, user1 *User) (*Comment, error) {
	setter := &CommentSetter{
		UserID: omit.From(user1.ID),
	}

	err := comment0.Update(ctx, exec, setter)
	if err != nil {
		return nil, fmt.Errorf("attachCommentUser0: %w", err)
	}

	return comment0, nil
}

func (comment0 *Comment) InsertUser(ctx context.Context, exec bob.Executor, related *UserSetter) error {
	user1, err := Users.Insert(related).One(ctx, exec)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}

	_, err = attachCommentUser0(ctx, exec, 1, comment0, user1)
	if err != nil {
		return err
	}

	comment0.R.User = user1

	user1.R.Comments = append(user1.R.Comments, comment0)

	return nil
}

func (comment0 *Comment) AttachUser(ctx context.Context, exec bob.Executor, user1 *User) error {
	var err error

	_, err = attachCommentUser0(ctx, exec, 1, comment0, user1)
	if err != nil {
		return err
	}

	comment0.R.User = user1

	user1.R.Comments = append(user1.R.Comments, comment0)

	return nil
}

// UpsertByPK uses an executor to upsert the Comment
func (s CommentSetter) UpsertByPK() bob.Mod[*dialect.InsertQuery] {
	pk := []string{
		"id",
	}

	conflictCols := []any{
		"id",
	}

	return im.OnConflict(conflictCols...).
		DoUpdate(im.SetExcluded(slices.DeleteFunc(s.SetColumns(), func(n string) bool {
			return slices.Contains(pk, n)
		})...))
}

// UpsertDoNothing uses an executor to upsert the Comment
func (s CommentSetter) UpsertDoNothing() bob.Mod[*dialect.InsertQuery] {
	return im.OnConflict().DoNothing()
}
