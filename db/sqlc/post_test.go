package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Sreekar-Panchangam/simplepost/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, user User) Post {
	arg := CreatePostParams{
		Title:  util.RandomTitle(),
		Body:   util.RandomBody(),
		UserID: int32(user.ID),
		Status: util.RandomStatus(),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Body, post.Body)
	require.Equal(t, arg.Status, post.Status)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	user := createRandomAccount(t)
	createRandomPost(t, user)
}

func TestGetPost(t *testing.T) {
	user := createRandomAccount(t)
	post1 := createRandomPost(t, user)
	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Body, post2.Body)
	require.Equal(t, post1.Status, post2.Status)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestListPost(t *testing.T) {
	for i := 0; i < 10; i++ {
		user := createRandomAccount(t)
		createRandomPost(t, user)
	}

	arg := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}

	posts, err := testQueries.ListPosts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, posts, 5)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}

func TestUpdatePost(t *testing.T) {
	user := createRandomAccount(t)
	post1 := createRandomPost(t, user)

	arg := UpdatePostParams{
		ID:   post1.ID,
		Body: util.RandomBody(),
	}

	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, arg.Body, post2.Body)
	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Status, post2.Status)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestDeletePost(t *testing.T) {
	user := createRandomAccount(t)
	post1 := createRandomPost(t, user)
	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}
