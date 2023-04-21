package commons

type ClaimsKey string

const (
	ClaimsKeyUserID  ClaimsKey = "user_id"
	ClaimsKeyIsAdmin ClaimsKey = "is_admin"
)

var (
	JwtKey = []byte("some_secret_key")
)
