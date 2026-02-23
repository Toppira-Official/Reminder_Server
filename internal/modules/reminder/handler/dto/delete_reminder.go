package dto

type DeleteReminderInput struct {
	ID uint `uri:"id" binding:"required,min=1"`
}
