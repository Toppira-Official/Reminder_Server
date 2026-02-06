package constants

type ReminderStatus string

const (
	Pending   ReminderStatus = "pending"
	Completed ReminderStatus = "completed"
	Missed    ReminderStatus = "missed"
)

var ReminderStatuses = [...]ReminderStatus{Pending, Completed, Missed}
