package DateDiff

import (
	"errors"
	"fmt"
)

// AbsoluteDateDifference computes the absolute difference in days between the two dates (x, y: unparsed as strings).
// Both x, y are expected to be of the following format: YYYY-MM-DD. The start and end day should not be counted and
// only the absolute difference in dates is considered. If the order of two dates being compared is flipped the
// difference remains the same (i.e.  AbsoluteDateDifference(x, y) == AbsoluteDateDifference(y, x)
func AbsoluteDateDifference(start, end string) (uint, error) {
	if len(start) != 10 || len(end) != 10 {
		return 0, errors.New(fmt.Sprintf(
			"dates must be of length 10. Start and end have lengths, %d and %d respectively", len(start), len(end),
		))
	}

	if !isCorrectStructure(start) {
		return 0, errors.New("start date has invalid structure. Expected YYYY-MM-DD")
	}
	if !isCorrectStructure(end) {
		return 0, errors.New("end date has invalid structure. Expected YYYY-MM-DD")
	}

	dateStart, err := constructDate(
		start[0:4],
		start[5:7],
		start[8:10],
	)
	if err != nil {
		return 0, fmt.Errorf("start date was invalid: %w", err)
	}

	dateEnd, err := constructDate(
		end[0:4],
		end[5:7],
		end[8:10],
	)
	if err != nil {
		return 0, fmt.Errorf("end date was invalid: %w", err)
	}

	return dateStart.AbsoluteDifference(dateEnd), nil
}

// isCorrectStructure checks if the date string has the correct structure for the date format YYYY-MM-DD. Equivalent to
// a string having the following regex pattern, ^.{4}-.{2}-.{2}$.
func isCorrectStructure(date string) bool {
	return len(date) == 10 && date[4] == '-' && date[7] == '-'
}

func constructDate(year, month, day string) (Date, error) {
	return Date{}, nil
}

// Date from the Gregorian calendar. Date structure are valid before created.
type Date struct {
	year  uint
	month uint
	day   uint
}

func (d Date) AbsoluteDifference(e Date) uint {
	return 0
}
