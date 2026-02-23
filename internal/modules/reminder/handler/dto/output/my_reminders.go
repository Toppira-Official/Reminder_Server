package output

import "github.com/Toppira-Official/Reminder_Server/internal/shared/entities"

type MyRemindersOutput struct {
	Reminders []*entities.Reminder `json:"reminders"`
} //	@name	MyRemindersOutput
