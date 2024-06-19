package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type Job struct {
	ID               string `validate:"uuid"`
	OutputBucketPath string `validate:"required"`
	Status           string `validate:"required"`
	Video            *Video
	VideoID          string
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewJob(output, status string, video *Video) (*Job,error) {
	j := &Job{OutputBucketPath: output, Status: status, Video: video, VideoID: video.ID}
	j.prepare()
	
	err := j.Validate()
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (j *Job) prepare() {
	j.ID = uuid.NewV4().String()
	j.CreatedAt = time.Now()
	j.UpdatedAt = time.Now()
}

func (j *Job) Validate() error {
	validator := validator.New()

	err := validator.Struct(j)

	return err
}
