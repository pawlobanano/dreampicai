package types

// userKey is a custom type for context key.
type userKey string

const UserContextKey userKey = "user"

// AuthenticatedUser is a struct that holds the user's email and login status.
type AuthenticatedUser struct {
	Email      string
	IsLoggedIn bool
}
