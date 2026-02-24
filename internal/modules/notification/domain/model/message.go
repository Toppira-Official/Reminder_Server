package model

type Message struct {
	Title    string  `json:"title"`
	Body     string  `json:"body"`
	Token    *string `json:"token,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
}
