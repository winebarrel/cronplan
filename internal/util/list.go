package util

import (
	"errors"
	"time"
)

func List(start int, end int, origin int, max int) ([]int, error) {
	if start < origin {
		return nil, errors.New("'start' must be >= 'origin'")
	} else if end < origin {
		return nil, errors.New("'end' must be >= 'origin'")
	} else if start > max {
		return nil, errors.New("'start' must be <= 'max'")
	} else if end > max {
		return nil, errors.New("'end' must be <= 'max'")
	}

	list := []int{}

	if start <= end {
		for i := start; i <= end; i++ {
			list = append(list, i)
		}
	} else {
		for i := start; i <= max; i++ {
			list = append(list, i)
		}

		for i := origin; i <= end; i++ {
			list = append(list, i)
		}
	}

	return list, nil
}

func ListMinute(start int, end int) ([]int, error) {
	return List(start, end, 0, 59)
}

func ListHour(start int, end int) ([]int, error) {
	return List(start, end, 0, 23)
}

func ListDayOfMonth(t time.Time, start int, end int) ([]int, error) {
	lom := LastOfMonth(t)

	if end > lom {
		end = lom
	}

	return List(start, end, 1, LastOfMonth(t))
}

func ListMonth(start time.Month, end time.Month) ([]time.Month, error) {
	nums, err := List(int(start), int(end), int(time.January), int(time.December))

	if err != nil {
		return nil, err
	}

	months := make([]time.Month, 0, len(nums))

	for _, n := range nums {
		months = append(months, time.Month(n))
	}

	return months, nil
}

func ListWeekday(start time.Weekday, end time.Weekday) ([]time.Weekday, error) {
	nums, err := List(int(start), int(end), int(time.Sunday), int(time.Saturday))

	if err != nil {
		return nil, err
	}

	wdays := make([]time.Weekday, 0, len(nums))

	for _, n := range nums {
		wdays = append(wdays, time.Weekday(n))
	}

	return wdays, nil
}

func ListYear(start int, end int) ([]int, error) {
	return List(start, end, 1970, 2199)
}
