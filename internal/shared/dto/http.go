package dto

type HttpOutput[T any] struct {
	Data T `json:"data,omitempty"`
} //	@name	HttpOutput
