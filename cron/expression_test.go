package cron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_genericHandler_Asterisk_OK(t *testing.T) {
	// GIVEN
	exp := "*"
	expectedValue := []int{3, 4, 5, 6, 7, 8}
	// WHEN
	values, err := genericHandler(exp, 3, 8)
	// THEN
	assert.Nil(t, err, "ERROR should be nil")
	if assert.NotNil(t, values) {
		assert.Equal(t, expectedValue, values, "they should be equal")
	}
}

func Test_genericHandler_Slash_OK(t *testing.T) {
	// GIVEN
	exp := "*/5"
	expectedValue := []int{3, 8, 13, 18}
	// WHEN
	values, err := genericHandler(exp, 3, 22)
	// THEN
	assert.Nil(t, err, "ERROR should be nil")
	if assert.NotNil(t, values) {
		assert.Equal(t, expectedValue, values, "they should be equal")
	}
}

func Test_genericHandler_Comma_OK(t *testing.T) {
	// GIVEN
	exp := "3,8,13,18"
	expectedValue := []int{3, 8, 13, 18}
	// WHEN
	values, err := genericHandler(exp, 3, 22)
	// THEN
	assert.Nil(t, err, "ERROR should be nil")
	if assert.NotNil(t, values) {
		assert.Equal(t, expectedValue, values, "they should be equal")
	}
}

func Test_genericHandler_Hyphen_OK(t *testing.T) {
	// GIVEN
	exp := "5-11"
	expectedValue := []int{5, 6, 7, 8, 9, 10, 11}
	// WHEN
	values, err := genericHandler(exp, 3, 22)
	// THEN
	assert.Nil(t, err, "ERROR should be nil")
	if assert.NotNil(t, values) {
		assert.Equal(t, expectedValue, values, "they should be equal")
	}
}

func Test_genericHandler_Format_KO(t *testing.T) {
	// GIVEN
	exp := "5/15-87"
	errExpected := fmt.Errorf("Can't handle %s ", exp)
	// WHEN
	_, err := genericHandler(exp, 3, 22)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_genericHandler_Format_KO2(t *testing.T) {
	// GIVEN
	exp := "15-17-18"
	errExpected := fmt.Errorf("Can't handle %s ", exp)
	// WHEN
	_, err := genericHandler(exp, 3, 22)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_genericHandler_Format_KO3(t *testing.T) {
	// GIVEN
	exp := "15/Hdheh"
	errExpected := fmt.Errorf("Can't handle %s ", exp)
	// WHEN
	_, err := genericHandler(exp, 3, 22)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_genericHandler_Hyphen_OutOfBounds_KO(t *testing.T) {
	// GIVEN 1
	exp := "15-87"
	errExpected := fmt.Errorf("Can't handle %s ", exp)
	// WHEN
	_, err := genericHandler(exp, 3, 22)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_genericHandler_Comma_OutOfBounds_KO(t *testing.T) {
	// GIVEN
	exp := "15,87"
	errExpected := fmt.Errorf("Can't handle %s ", "87")
	// WHEN
	_, err := genericHandler(exp, 3, 22)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_minutesExtract_OK(t *testing.T) {} //TODO : but error test is useful (for me)

func Test_hoursExtract_OK(t *testing.T) {} //TODO : but error test is useful (for me)

func Test_daysOfMonthExtract_OK(t *testing.T) {} //TODO : but error test is useful (for me)

func Test_monthExtract_OK(t *testing.T) {} //TODO : but error test is useful (for me)

func Test_daysOfWeekExtract_OK(t *testing.T) {} //TODO : but error test is useful (for me)

//---- Test OutOfBounds ---
func Test_minutesExtract_OutOfBounds_KO(t *testing.T) {
	// GIVEN
	expString := "0-67"
	exp := GetExpression("myCommand")
	errExpected := fmt.Errorf("Can't handle %s ", expString)
	// WHEN
	err := exp.minutesExtract(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_hoursExtract_OutOfBounds_KO(t *testing.T) {
	// GIVEN
	expString := "0-25"
	exp := GetExpression("myCommand")
	errExpected := fmt.Errorf("OutOfBounds %s - min : 0 - max : 23", expString)
	// WHEN
	err := exp.hoursExtract(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_monthExtract_OutOfBounds_KO(t *testing.T) {
	// GIVEN
	expString := "0-13"
	exp := GetExpression("myCommand")
	errExpected := fmt.Errorf("OutOfBounds %s - min : 1 - max : 12", expString)
	// WHEN
	err := exp.monthExtract(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_daysOfMonthExtract_OutOfBounds_KO(t *testing.T) {
	// GIVEN
	expString := "0-33"
	exp := GetExpression("myCommand")
	errExpected := fmt.Errorf("OutOfBounds %s - min : 1 - max : 31", expString)
	// WHEN
	err := exp.daysOfMonthExtract(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}

func Test_daysOfWeekExtract_OutOfBounds_KO(t *testing.T) {
	// GIVEN
	expString := "0-8"
	exp := GetExpression("myCommand")
	errExpected := fmt.Errorf("OutOfBounds %s - min : 1 - max : 7", expString)
	// WHEN
	err := exp.daysOfWeekExtract(expString)
	// THEN
	assert.NotNil(t, err, "ERROR should happen")
	assert.Equal(t, err, errExpected, "ERROR should happen")
}
