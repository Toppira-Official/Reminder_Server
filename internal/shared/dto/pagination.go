package dto

type PaginationInput struct {
	Page  int `form:"page,default=1" binding:"min=1"`
	Limit int `form:"limit,default=10" binding:"min=5,max=20"`
} //	@name	PaginationInput
