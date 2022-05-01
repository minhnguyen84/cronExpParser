package cron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Parse_OK(t *testing.T) {
	// GIVEN
	expString := "*/15 0 1,15 * 1-5 /usr/bin/find"
	expectedExp := Expression{
		command:     "/usr/bin/find",
		minutes:     []int{0, 15, 30, 45},
		hours:       []int{0},
		daysOfMonth: []int{1, 15},
		months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		daysOfWeek:  []int{1, 2, 3, 4, 5},
	}
	// WHEN
	values, err := Parse(expString)
	// THEN
	assert.Nil(t, err, "ERROR should be nil")
	if assert.NotNil(t, values) {
		assert.Equal(t, &expectedExp, values, "they should be equal")
	}
}

func Test_Parse_TooMuchArg_K0(t *testing.T) {
	// GIVEN
	expString := "*/15 0 1,15 * 1-5 98 /usr/bin/find"
	errExpected := fmt.Errorf("Too much argument %s ", expString)
	// WHEN
	_, err := Parse(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should be nil")
	assert.Equal(t, errExpected, err, "they should be equal")
}

func Test_Parse_TooMuchArg_K02(t *testing.T) {
	// GIVEN
	expString := "*/15 0 1, 15 * 1-5 /usr/bin/find"
	errExpected := fmt.Errorf("Too much argument %s ", expString)
	// WHEN
	_, err := Parse(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should be nil")
	assert.Equal(t, errExpected, err, "they should be equal")
}

func Test_Parse_NotEnoughArg_K0(t *testing.T) {
	// GIVEN
	expString := "*/15 0 1,15 * /usr/bin/find"
	errExpected := fmt.Errorf("Not enough argument %s ", expString)
	// WHEN
	_, err := Parse(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should be nil")
	assert.Equal(t, errExpected, err, "they should be equal")
}
