package main

import (
	"context"
	"database/sql"
	"fmt"
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
			Email: omit.From(fmt.Sprintf("email-%d@omerid.com", i)),
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

	if _, err := models.Users.Insert(userInsertQueries).All(ctx, client); err != nil {
		t.Fatalf("failed to insert user %v", err)
	}

	if _, err := models.Posts.Insert(postInsertQueries).All(ctx, client); err != nil {
		t.Fatalf("failed to insert post %v", err)
	}

	t.Run("ThenLoad", func(t *testing.T) {
		var users models.UserSlice
		if users, err = models.Users.Query(
			// psql.WhereOr(
			// 	models.SelectWhere.Users.ID.EQ(1),
			// 	models.SelectWhere.Users.ID.EQ(2),
			// ),
			models.ThenLoadUserPosts(),
		).All(ctx, bob.Debug(client)); err != nil {
			t.Fatalf("Failed to reload user: %v", err)
		}
		_ = users
		// spew.Dump(users)
	})
}
