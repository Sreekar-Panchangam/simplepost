package util

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomRole() string {
	roles := []string{"User", "Admin", "Owner"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

func RandomTitle() string {
	return RandomString(6)
}

func RandomBody() sql.NullString {
	body := RandomString(30)
	return sql.NullString{String: body, Valid: true}
}

func RandomStatus() string {
	stat := []string{"Posted", "Draft"}
	n := len(stat)
	return stat[rand.Intn(n)]
}
