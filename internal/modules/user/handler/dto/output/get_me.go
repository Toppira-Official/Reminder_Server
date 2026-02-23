package output

import sharedDto "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"

type GetMeOutput struct {
	User *sharedDto.UserOutput `json:"user"`
} //	@name	GetMeOutput
