// Code generated by sadayuki-matsuno bob-ex. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"hash/maphash"

	"github.com/lib/pq"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/clause"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

var TableNames = struct {
	Comments string
	PostTags string
	Posts    string
	Tags     string
	Users    string
}{
	Comments: "comments",
	PostTags: "post_tags",
	Posts:    "posts",
	Tags:     "tags",
	Users:    "users",
}

var ColumnNames = struct {
	Comments commentColumnNames
	PostTags postTagColumnNames
	Posts    postColumnNames
	Tags     tagColumnNames
	Users    userColumnNames
}{
	Comments: commentColumnNames{
		ID:        "id",
		PostID:    "post_id",
		UserID:    "user_id",
		Body:      "body",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	PostTags: postTagColumnNames{
		PostID:    "post_id",
		TagID:     "tag_id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	Posts: postColumnNames{
		ID:        "id",
		UserID:    "user_id",
		Title:     "title",
		Content:   "content",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	Tags: tagColumnNames{
		ID:        "id",
		Name:      "name",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	Users: userColumnNames{
		ID:        "id",
		Name:      "name",
		Email:     "email",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
}

var (
	SelectWhere     = Where[*dialect.SelectQuery]()
	UpdateWhere     = Where[*dialect.UpdateQuery]()
	DeleteWhere     = Where[*dialect.DeleteQuery]()
	OnConflictWhere = Where[*clause.ConflictClause]() // Used in ON CONFLICT DO UPDATE
)

func Where[Q psql.Filterable]() struct {
	Comments commentWhere[Q]
	PostTags postTagWhere[Q]
	Posts    postWhere[Q]
	Tags     tagWhere[Q]
	Users    userWhere[Q]
} {
	return struct {
		Comments commentWhere[Q]
		PostTags postTagWhere[Q]
		Posts    postWhere[Q]
		Tags     tagWhere[Q]
		Users    userWhere[Q]
	}{
		Comments: buildCommentWhere[Q](CommentColumns),
		PostTags: buildPostTagWhere[Q](PostTagColumns),
		Posts:    buildPostWhere[Q](PostColumns),
		Tags:     buildTagWhere[Q](TagColumns),
		Users:    buildUserWhere[Q](UserColumns),
	}
}

var (
	SelectJoins = getJoins[*dialect.SelectQuery]()
	UpdateJoins = getJoins[*dialect.UpdateQuery]()
	DeleteJoins = getJoins[*dialect.DeleteQuery]()
)

type joinSet[Q interface{ aliasedAs(string) Q }] struct {
	InnerJoin Q
	LeftJoin  Q
	RightJoin Q
}

func (j joinSet[Q]) AliasedAs(alias string) joinSet[Q] {
	return joinSet[Q]{
		InnerJoin: j.InnerJoin.aliasedAs(alias),
		LeftJoin:  j.LeftJoin.aliasedAs(alias),
		RightJoin: j.RightJoin.aliasedAs(alias),
	}
}

type joins[Q dialect.Joinable] struct {
	Comments joinSet[commentJoins[Q]]
	PostTags joinSet[postTagJoins[Q]]
	Posts    joinSet[postJoins[Q]]
	Tags     joinSet[tagJoins[Q]]
	Users    joinSet[userJoins[Q]]
}

func buildJoinSet[Q interface{ aliasedAs(string) Q }, C any, F func(C, string) Q](c C, f F) joinSet[Q] {
	return joinSet[Q]{
		InnerJoin: f(c, clause.InnerJoin),
		LeftJoin:  f(c, clause.LeftJoin),
		RightJoin: f(c, clause.RightJoin),
	}
}

func getJoins[Q dialect.Joinable]() joins[Q] {
	return joins[Q]{
		Comments: buildJoinSet[commentJoins[Q]](CommentColumns, buildCommentJoins),
		PostTags: buildJoinSet[postTagJoins[Q]](PostTagColumns, buildPostTagJoins),
		Posts:    buildJoinSet[postJoins[Q]](PostColumns, buildPostJoins),
		Tags:     buildJoinSet[tagJoins[Q]](TagColumns, buildTagJoins),
		Users:    buildJoinSet[userJoins[Q]](UserColumns, buildUserJoins),
	}
}

type modAs[Q any, C interface{ AliasedAs(string) C }] struct {
	c C
	f func(C) bob.Mod[Q]
}

func (m modAs[Q, C]) Apply(q Q) {
	m.f(m.c).Apply(q)
}

func (m modAs[Q, C]) AliasedAs(alias string) bob.Mod[Q] {
	m.c = m.c.AliasedAs(alias)
	return m
}

func randInt() int64 {
	out := int64(new(maphash.Hash).Sum64())

	if out < 0 {
		return -out % 10000
	}

	return out % 10000
}

// ErrUniqueConstraint captures all unique constraint errors by explicitly leaving `s` empty.
var ErrUniqueConstraint = &UniqueConstraintError{s: ""}

type UniqueConstraintError struct {
	// s is a string uniquely identifying the constraint in the raw error message returned from the database.
	s string
}

func (e *UniqueConstraintError) Error() string {
	return e.s
}

func (e *UniqueConstraintError) Is(target error) bool {
	err, ok := target.(*pq.Error)
	if !ok {
		return false
	}
	return err.Code == "23505" && (e.s == "" || err.Constraint == e.s)
}
