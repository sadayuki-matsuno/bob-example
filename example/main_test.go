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
	"github.com/davecgh/go-spew/spew"
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
	var commentInsertQueries bob.Mods[*dialect.InsertQuery]
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

		comment1 := &models.CommentSetter{
			PostID: omit.From(int32(i)),
			UserID: omit.From(int32(i)),
			Body:   omit.From(fmt.Sprintf("Comment body 1 %d", i)),
		}
		comment2 := &models.CommentSetter{
			PostID: omit.From(int32(i)),
			UserID: omit.From(int32(i)),
			Body:   omit.From(fmt.Sprintf("Comment body 2 %d", i)),
		}
		commentInsertQueries = append(commentInsertQueries, comment1, comment2)
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

	for commentInsertQuery := range slices.Chunk(commentInsertQueries, 10000) {
		if _, err := models.Comments.Insert(commentInsertQuery).All(ctx, client); err != nil {
			t.Fatalf("failed to insert comment %v", err)
		}
	}

	t.Run("ThenLoad", func(t *testing.T) {
		var users models.UserSlice
		if users, err = models.Users.Query(
			models.SelectWhere.Users.ID.In([]int32{2, 3, 4}...),
			models.ThenLoadUserPosts(
				models.SelectWhere.Posts.ID.In([]int32{1, 2, 3, 4}...),
				models.ThenLoadPostComments(),
				models.PreloadPostUser(
					models.ThenLoadUserPosts(),
				),
			),
		).All(ctx, bob.Debug(client)); err != nil {
			t.Fatalf("Failed to reload user: %v", err)
		}

		for _, user := range users {
			for _, post := range user.R.Posts {
				for _, comment := range post.R.Comments {
					spew.Dump(comment)
				}
				for _, postUser := range post.R.User.R.Posts {
					spew.Dump(postUser)
				}
			}
			spew.Dump(user)
		}
	})
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	sqldb, err := sql.Open("txdb", fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	if err != nil {
		t.Fatal(err)
	}
	defer sqldb.Close()
	client := bob.NewDB(sqldb)

	user := &models.UserSetter{
		Name:  omit.From("User"),
		Email: omit.From("email@example.com"),
	}
	var newUser *models.User
	if newUser, err = models.Users.Insert(user).One(ctx, client); err != nil {
		t.Fatalf("failed to insert user %v", err)
	}

	var newUser2 *models.User
	if newUser2, err = models.Users.Update(
		models.UpdateWhere.Users.ID.EQ(newUser.ID),
		(models.UserSetter{
			Email: omit.From("new_email@example.com"),
		}).UpdateMod(),
	).One(ctx, client); err != nil {
		t.Fatalf("failed to update user %v", err)
	}
	spew.Dump(newUser2)
}

func TestUpsert(t *testing.T) {
	ctx := context.Background()
	sqldb, err := sql.Open("txdb", fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	if err != nil {
		t.Fatal(err)
	}
	defer sqldb.Close()
	client := bob.NewDB(sqldb)

	user1 := &models.UserSetter{
		ID:    omit.From(int32(1)),
		Name:  omit.From("User"),
		Email: omit.From("email@example.com"),
	}

	if _, err = models.Users.Insert(
		user1,
	).Exec(ctx, client); err != nil {
		t.Fatalf("failed to insert user %v", err)
	}

	user1update := &models.UserSetter{
		Name:  omit.From("User2"),
		Email: omit.From("email1update@example.com"),
	}

	user2 := &models.UserSetter{
		Name:  omit.From("User3"),
		Email: omit.From("email@example.com"),
	}

	if _, err = models.Users.Insert(
		user1update,
		user2,
		user1.UpsertByUsersEmailKey(),
	).Exec(ctx, bob.Debug(client)); err != nil {
		t.Fatalf("failed to insert user %v", err)
	}

	var allUsers models.UserSlice
	if allUsers, err = models.Users.Query().All(ctx, client); err != nil {
		t.Fatalf("failed to select user %v", err)
	}
	spew.Dump(allUsers)
}
