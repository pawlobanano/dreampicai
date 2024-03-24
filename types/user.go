package types

type userKey string

const UserContextKey userKey = "user"

type AuthenticatedUser struct {
	Email      string
	IsLoggedIn bool
}
