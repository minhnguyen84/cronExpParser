package cron

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	fieldFinder = regexp.MustCompile(`\S+`)
	entryFinder = regexp.MustCompile(`[^,]+`)
	rangeMask   = regexp.MustCompile(`^(0?[0-9]|[1-5][0-9])-(0?[0-9]|[1-5][0-9])$`)
	stepMask    = regexp.MustCompile(`^(0?[0-9]|[1-5][0-9]|\*)/(0?[0-9]|[1-5][0-9])$`)
	numberMask  = regexp.MustCompile(`^(0?[0-9]|[1-5][0-9])$`)
)

type Expression struct {
	command     string
	minutes     []int
	hours       []int
	daysOfMonth []int
	months      []int
	daysOfWeek  []int
}

func GetExpression(command string) *Expression {
	return &Expression{command: command}
}

func (expr *Expression) minutesExtract(s string) error {
	values, err := genericHandler(s, 0, 59)
	if err != nil {
		return err
	}
	expr.minutes = values
	return nil
}

func (expr *Expression) hoursExtract(s string) error {
	values, err := genericHandler(s, 0, 23)
	if err != nil {
		return err
	}
	expr.hours = values
	return nil
}

func (expr *Expression) daysOfMonthExtract(s string) error {
	values, err := genericHandler(s, 1, 31)
	if err != nil {
		return err
	}
	expr.daysOfMonth = values
	return nil
}

func (expr *Expression) monthExtract(s string) error {
	values, err := genericHandler(s, 1, 12)
	if err != nil {
		return err
	}
	expr.months = values
	return nil
}

func (expr *Expression) daysOfWeekExtract(s string) error {
	values, err := genericHandler(s, 1, 7)
	if err != nil {
		return err
	}
	expr.daysOfWeek = values
	return nil
}

func (expr *Expression) ToString() []string {
	result := make([]string, 6)
	result[0] = formatString("minute", expr.minutes)
	result[1] = formatString("hour", expr.hours)
	result[2] = formatString("day of month", expr.daysOfMonth)
	result[3] = formatString("month", expr.months)
	result[4] = formatString("day of week", expr.daysOfWeek)
	result[5] = fmt.Sprintf("command       %s", expr.command) //TODO : could do it better

	return result
}

func genericHandler(s string, min, max int) ([]int, error) {
	//all
	if s == "*" {
		return makeRange(min, max), nil
	}

	//manage the complexes cases
	//cutout "," -> "1,5-6" => ["1", "5-6"]
	indices := entryFinder.FindAllStringIndex(s, -1)

	if len(indices) == 0 {
		return nil, fmt.Errorf("No entry %s ", s)
	}
	result := make([]int, 0)
	for _, entry := range indices {
		entryVal := s[entry[0]:entry[1]]
		// single value "5"
		if numberMask.MatchString(entryVal) {
			i, err := strconv.Atoi(entryVal)
			if err != nil {
				return nil, err
			}
			result = append(result, i)
			continue
		}
		// range "5-15"
		if rangeMask.MatchString(entryVal) {
			vals := strings.Split(entryVal, "-")
			if len(vals) != 2 {
				// we can't handle it
				return nil, fmt.Errorf("Can't handle %s ", entryVal)
			}
			minTmp, err := strconv.Atoi(vals[0])
			if err != nil {
				return nil, err
			}
			maxTmp, err := strconv.Atoi(vals[1])
			if err != nil {
				return nil, err
			}
			if minTmp > maxTmp || minTmp < min || maxTmp > max {
				// we can't handle it
				return nil, fmt.Errorf("OutOfBounds %s - min : %d - max : %d", entryVal, min, max)
			}
			result = append(result, makeRange(minTmp, maxTmp)...)
			continue
		}

		// step "0/15" -> normalise "*/15" to "{min}/15"
		if stepMask.MatchString(entryVal) {
			entryValTmp := strings.ReplaceAll(entryVal, "*", strconv.Itoa(min))
			vals := strings.Split(entryValTmp, "/")
			firstVal, err := strconv.Atoi(vals[0])
			if err != nil {
				return nil, err
			}
			stepVal, err := strconv.Atoi(vals[1])
			if err != nil {
				return nil, err
			}

			if firstVal > max {
				// we can't handle it
				return nil, fmt.Errorf("OutOfBounds %s : firstVal is big than max ", entryVal)
			}
			val := firstVal
			for i := 1; val <= max; i++ {
				result = append(result, val)
				val += stepVal
			}
			continue
		}
		// we can't handle it
		return nil, fmt.Errorf("Can't handle %s ", entryVal)
	}
	return result, nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func formatString(title string, values []int) string {
	v := title
	for i := len(title); i < 13; i++ {
		v += " "
	}

	for _, value := range values {
		v += fmt.Sprintf(" %d", value)
	}

	return v
}
