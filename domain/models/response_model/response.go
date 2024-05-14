package response_model

type Response struct {
	Status  string      `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data,omitempty"`
}
