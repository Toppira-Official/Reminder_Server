package entities

type User struct {
	Base

	Email string  `json:"email"`
	Phone *string `json:"phone,omitempty"`

	Name           *string `json:"name,omitempty"`
	ProfilePicture *string `json:"profile_picture,omitempty"`

	Password *string `json:"-"`
}
