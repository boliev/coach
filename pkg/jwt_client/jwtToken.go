package jwt_client

type JwtToken struct {
	Claims map[string]interface{}
	Valid  bool
}
