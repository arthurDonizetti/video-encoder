package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	uuid "github.com/satori/go.uuid"

	"github.com/arthurDonizetti/video-encoder/encoder/internal/domain"
)

func TestVideo_Validate(t *testing.T) {
	type args struct {
		video *domain.Video
	}
	type want struct {
		err error
	}

	tests := []struct {
		name string
		args
		want
	}{
		{
			name: "Video is empty",
			args: args{video: domain.NewVideo()},
			want: want{err: assert.AnError},
		},
		{
			name: "Video ID is not an uuid",
			args: args{video: &domain.Video{ID: "wrong_format", ResourceID: "any", FilePath: "any"}},
			want: want{err: assert.AnError},
		},
		{
			name: "Video ID is an uuid",
			args: args{video: &domain.Video{ID: uuid.NewV4().String(), ResourceID: "any", FilePath: "any"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.video.Validate()

			if tt.want.err != nil {
				assert.Error(t, err)
			}
		})
	}
}
