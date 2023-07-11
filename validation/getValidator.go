package validation

type GetRequested struct {
	Key string `json:"key" binding:"required"`
}
