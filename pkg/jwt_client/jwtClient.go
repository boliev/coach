package jwt_client

type JwtClient interface {
	Parse(tokenString string) (*JwtToken, error)
}
