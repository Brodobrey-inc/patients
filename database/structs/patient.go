package structs

import "github.com/google/uuid"

type Patient struct {
	FullName string    `mapstructure:"fullname"`
	Birthday string    `mapstructure:"birthday"`
	Gender   int       `mapstructure:"gender"`
	GUID     uuid.UUID `mapstructure:"guid"`
}
