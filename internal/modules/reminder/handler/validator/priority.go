package validator

import (
	"reflect"
	"slices"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/constants"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterPriorityValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("priority", func(fl validator.FieldLevel) bool {
			field := fl.Field()

			if field.Kind() == reflect.Ptr {
				if field.IsNil() {
					return true
				}
				field = field.Elem()
			}

			val, ok := field.Interface().(constants.ReminderPriority)
			if !ok {
				return false
			}

			return slices.Contains(constants.ReminderPriorities[:], val)
		})
		if err != nil {
			return
		}
	}
}
