package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/arthurDonizetti/video-encoder/encoder/internal/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestJob_New(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("outputPath", "converted", video)

	fmt.Println(err)

	assert.NotNil(t, job)
	assert.Nil(t, err)
}
