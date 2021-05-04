package response

type badRequest struct {
	Error string `json:"error"`
}

func BadRequest(error string) *badRequest {
	return &badRequest{
		Error: error,
	}
}
