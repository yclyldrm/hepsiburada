package test

import (
	"fmt"
	"hbcase/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSystemTime(t *testing.T) {
	systemTime := utils.GetSystemTime()

	expected := utils.START_TIME_FORMAT

	assert.Equal(t, expected, systemTime)
}

func TestIncreaseTimeWithSuccess(t *testing.T) {
	utils.SetStartTime()

	err := utils.IncreaseTime(1)

	assert.NoError(t, err)

	actual := utils.GetSystemTime()
	expected := "01:00"

	assert.Equal(t, expected, actual, fmt.Sprintf("expected time:%s , actual time:%s", expected, actual))
}

func TestIncreaseTimeWithFail(t *testing.T) {
	utils.SetStartTime()

	err := utils.IncreaseTime(1)

	assert.NoError(t, err)

	actual := utils.GetSystemTime()
	expected := "02:00"

	assert.NotEqual(t, expected, actual, fmt.Sprintf("expected time:%s , actual time:%s", expected, actual))
}

func TestTimeDifference(t *testing.T) {
	utils.SetStartTime()

	actual := utils.TimeDifference("06:00")
	expected := 6

	assert.Equal(t, expected, actual, fmt.Sprintf("expected time diff:%d, actual time diff:%d", expected, actual))
}
