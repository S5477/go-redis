package models

type Add struct {
	Key   string `json:"key" binding:"required,min=3"`
	Value string `json:"value" binding:"required,min=3"`
}

type Get struct {
	Key string `form:"key" binding:"required,min=3"`
}
