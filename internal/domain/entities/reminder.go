package entities

import (
	"time"

	"github.com/Toppira-Official/backend/internal/domain/constants"
)

type Reminder struct {
	Base

	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`

	Status        constants.ReminderStatus `json:"status"`
	ReminderTimes []time.Time              `json:"reminder_times,omitempty"`
	ScheduledAt   time.Time                `json:"scheduled_at"`

	Priority *string `json:"priority"`

	UserID uint `json:"user_id"`
}
