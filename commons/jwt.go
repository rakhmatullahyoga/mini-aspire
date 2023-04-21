package commons

type ClaimsKey string

const (
	ClaimsKeyUserID  ClaimsKey = "user_id"
	ClaimsKeyIsAdmin ClaimsKey = "is_admin"
	JwtKey                     = "some_secret_key"
)
