package domain_test

import (
	"fullcycle-video-encoder/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}
