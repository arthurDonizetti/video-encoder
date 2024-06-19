package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Video struct {
	ID         string `validate:"uuid"`
	ResourceID string `validate:"required"`
	FilePath   string `validate:"required"`
	CreatedAt  time.Time
}

func NewVideo() *Video {
	return &Video{
		ID:         "",
		ResourceID: "",
		FilePath:   "",
		CreatedAt:  time.Time{},
	}
}

func (v *Video) Validate() error {
	validator := validator.New()

	err := validator.Struct(v)

	if err != nil {
		return err
	}

	return nil
}
