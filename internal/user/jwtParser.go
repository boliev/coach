package user

type JwtParser interface {
	Parse(tokenString string) (string, error)
}
