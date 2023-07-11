package validation

type AddRequested struct {
	Key   string `json:"key" binding:"required,min=3"`
	Value string `json:"value" binding:"required,min=3"`
}
