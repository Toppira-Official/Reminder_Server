package input

type DeleteReminderInput struct {
	ID uint `uri:"id" binding:"required,min=1"`
}
