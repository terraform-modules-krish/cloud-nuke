package commands

import (
	"testing"
	"time"

	gruntwork-cli "github.com/terraform-modules-krish/go-commons/errors"
	"github.com/stretchr/testify/assert"
)

func TestParseDuration(t *testing.T) {
	now := time.Now()
	then, err := parseDurationParam("1h")
	if err != nil {
		assert.Fail(t, errors.WithStackTrace(err).Error())
	}

	if now.Hour() == 0 {
		// At midnight, now.Hour returns 0 so we need to handle that specially.
		assert.Equal(t, 23, then.Hour())
		// Also, the date changed, so 1 hour ago will be the previous day.
		assert.Equal(t, now.Day()-1, then.Day())
	} else {
		assert.Equal(t, now.Hour()-1, then.Hour())
		assert.Equal(t, now.Day(), then.Day())
	}

	assert.Equal(t, now.Month(), then.Month())
	assert.Equal(t, now.Year(), then.Year())
}

func TestParseDurationInvalidFormat(t *testing.T) {
	_, err := parseDurationParam("")
	assert.Error(t, err)
}
