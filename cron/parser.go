package cron

import "fmt"

func Parse(sExpression string) (*Expression, error) {
	indices := fieldFinder.FindAllStringIndex(sExpression, -1)

	if len(indices) < 6 {
		return nil, fmt.Errorf("Not enough argument %s ", sExpression)
	}

	if len(indices) > 6 {
		return nil, fmt.Errorf("Too much argument %s ", sExpression)
	}

	exp := GetExpression(sExpression[indices[5][0]:indices[5][1]])

	i := 0
	err := exp.minutesExtract(sExpression[indices[i][0]:indices[i][1]])
	if err != nil {
		return nil, fmt.Errorf("minutesExtract - %w", err)
	}
	i++

	err = exp.hoursExtract(sExpression[indices[i][0]:indices[i][1]])
	if err != nil {
		return nil, fmt.Errorf("hoursExtract - %w", err)
	}
	i++

	err = exp.daysOfMonthExtract(sExpression[indices[i][0]:indices[i][1]])
	if err != nil {
		return nil, fmt.Errorf("daysOfMonthExtract - %w", err)
	}
	i++

	err = exp.monthExtract(sExpression[indices[i][0]:indices[i][1]])
	if err != nil {
		return nil, fmt.Errorf("monthExtract - %w", err)
	}
	i++

	err = exp.daysOfWeekExtract(sExpression[indices[i][0]:indices[i][1]])
	if err != nil {
		return nil, fmt.Errorf("daysOfWeekExtract - %w", err)
	}

	return exp, nil
}
