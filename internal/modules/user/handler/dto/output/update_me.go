package output

import sharedDto "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"

type UpdateMeOutput struct {
	User *sharedDto.UserOutput `json:"user"`
} //	@name	UpdateMeOutput
