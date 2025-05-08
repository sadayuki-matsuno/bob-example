package main

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/sadayuki-matsuno/bob-example/example/models"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

func init() {
	txdb.Register("txdb", "postgres", "postgres://bob:test@localhost:5433/testdb?sslmode=disable")
}

func TestTheLoad(t *testing.T) {
	ctx := context.Background()
	sqldb, err := sql.Open("txdb", fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	if err != nil {
		t.Fatal(err)
	}
	defer sqldb.Close()
	client := bob.NewDB(sqldb)

	var userInsertQueries bob.Mods[*dialect.InsertQuery]
	var postInsertQueries bob.Mods[*dialect.InsertQuery]
	for i := 1; i <= 100000; i++ {
		user := &models.UserSetter{
			ID:    omit.From(int32(i)),
			Name:  omit.From(fmt.Sprintf("User %d", i)),
			Email: omit.From(fmt.Sprintf("email-%d@example.com", i)),
			// 必要に応じて他のフィールドも指定
		}
		userInsertQueries = append(userInsertQueries, user)

		post := &models.PostSetter{
			ID:      omit.From(int32(i)),
			UserID:  omit.From(int32(i)),
			Title:   omit.From(fmt.Sprintf("Post %d", i)),
			Content: omitnull.From(fmt.Sprintf("Post body %d", i)),
		}
		postInsertQueries = append(postInsertQueries, post)
	}

	for userInsertQuery := range slices.Chunk(userInsertQueries, 10000) {
		if _, err := models.Users.Insert(userInsertQuery).All(ctx, client); err != nil {
			t.Fatalf("failed to insert user %v", err)
		}
	}

	for postInsertQuery := range slices.Chunk(postInsertQueries, 10000) {
		if _, err := models.Posts.Insert(postInsertQuery).All(ctx, client); err != nil {
			t.Fatalf("failed to insert post %v", err)
		}
	}

	t.Run("ThenLoad", func(t *testing.T) {
		var users models.UserSlice
		if users, err = models.Users.Query(
			models.SelectWhere.Users.ID.In([]int32{2, 3, 4}...),
			models.ThenLoadUserPosts(
				models.SelectWhere.Posts.ID.In([]int32{1, 2, 3, 4}...),
				models.PreloadPostUser(
					models.ThenLoadUserPosts(),
				),
			),
		).All(ctx, bob.Debug(client)); err != nil {
			t.Fatalf("Failed to reload user: %v", err)
		}
		_ = users
		// spew.Dump(users)
	})
}
