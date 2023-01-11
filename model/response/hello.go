package response

type HelloResponse struct {
	Message string `json:"message" form:"message"`
}
