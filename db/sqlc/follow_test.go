package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomFollow(t *testing.T, user1, user2 User) Follow {
	arg := CreateFollowParams{
		FollowingUserID: int32(user1.ID),
		FollowedUserID:  int32(user2.ID),
	}

	follow, err := testQueries.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)

	require.Equal(t, arg.FollowingUserID, follow.FollowingUserID)
	require.Equal(t, arg.FollowedUserID, follow.FollowedUserID)

	require.NotZero(t, follow.CreatedAt)

	return follow
}

func TestGetFollowing(t *testing.T) {
	user1 := createRandomAccount(t)
	user2 := createRandomAccount(t)
	follow1 := createRandomFollow(t, user1, user2)
	follow2, err := testQueries.GetFollowing(context.Background(), follow1.FollowingUserID)
	require.NoError(t, err)
	require.NotEmpty(t, follow2)

	require.Equal(t, follow1.FollowingUserID, follow2.FollowingUserID)
	require.Equal(t, follow1.FollowedUserID, follow2.FollowedUserID)
	require.WithinDuration(t, follow1.CreatedAt, follow2.CreatedAt, time.Second)
}

func TestGetFollower(t *testing.T) {
	user1 := createRandomAccount(t)
	user2 := createRandomAccount(t)
	follow1 := createRandomFollow(t, user1, user2)
	follow2, err := testQueries.GetFollower(context.Background(), follow1.FollowedUserID)
	require.NoError(t, err)
	require.NotEmpty(t, follow2)

	require.Equal(t, follow1.FollowingUserID, follow2.FollowingUserID)
	require.Equal(t, follow1.FollowedUserID, follow2.FollowedUserID)
	require.WithinDuration(t, follow1.CreatedAt, follow2.CreatedAt, time.Second)
}

func TestListFollowing(t *testing.T) {
	for i := 0; i < 10; i++ {
		user1 := createRandomAccount(t)
		user2 := createRandomAccount(t)
		createRandomFollow(t, user1, user2)
	}

	arg := ListFollowingParams{
		Limit:  5,
		Offset: 5,
	}

	follows, err := testQueries.ListFollowing(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, follows, 5)

	for _, follow := range follows {
		require.NotEmpty(t, follow)
	}
}

func TestListFollower(t *testing.T) {
	for i := 0; i < 10; i++ {
		user1 := createRandomAccount(t)
		user2 := createRandomAccount(t)
		createRandomFollow(t, user1, user2)
	}

	arg := ListFollowerParams{
		Limit:  5,
		Offset: 5,
	}

	follows, err := testQueries.ListFollower(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, follows, 5)

	for _, follow := range follows {
		require.NotEmpty(t, follow)
	}
}

func TestDeleteFollow(t *testing.T) {
	user1 := createRandomAccount(t)
	user2 := createRandomAccount(t)
	follow1 := createRandomFollow(t, user1, user2)
	err := testQueries.DeleteFollow(context.Background(), follow1.FollowingUserID)
	require.NoError(t, err)

	follow2, err := testQueries.GetFollowing(context.Background(), follow1.FollowingUserID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, follow2)
}
