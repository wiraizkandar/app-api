package response

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Constants for API response status
const (
	SuccessStatus = "success"
	ErrorStatus   = "error"
)
