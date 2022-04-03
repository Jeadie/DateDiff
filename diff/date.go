package diff

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

// constructDate validates the year, month, day parameters and constructs a Date if valid.
func constructDate(year, month, day string) (Date, error) {
	y, err := UintParse(year)
	if err != nil {
		return Date{}, fmt.Errorf("year is not a valid integer because %w", err)
	}
	m, err := UintParse(month)
	if err != nil {
		return Date{}, fmt.Errorf("month is not a valid integer because %w", err)
	}
	d, err := UintParse(day)
	if err != nil {
		return Date{}, fmt.Errorf("day is not a valid integer because %w", err)
	}

	// Validate if y, m, d construct a valid date.
	err = validateDate(y, m, d)
	if err != nil {
		return Date{}, fmt.Errorf("")
	}

	date := Date{
		year:  y,
		month: m,
		day:   d,
	}
	return date, nil
}

// validateDate throws an error if the date provided is not valid.
func validateDate(y, m, d uint) error {
	if m == 0 || m > 12 {
		return errors.New("month is not within range [1, 12]")
	}
	if isLeapYear(y) {
		if d == 0 || d > monthDatesLeap[m-1] {
			return fmt.Errorf("in month %d, there are not %d days during a leap year", m, d)
		}
	} else {
		if d == 0 || d > monthDates[m-1] {
			return fmt.Errorf("in month %d, there are not %d days during a non-leap year", m, d)
		}
	}
	return nil
}

// isLeapYear returns true if the year is a leap year.
func isLeapYear(year uint) bool {
	return year%4 == 0
}

// Date from the Gregorian calendar. Date structure are valid before created.
type Date struct {
	year  uint
	month uint
	day   uint
}

func (d Date) AbsoluteDifference(e Date) uint {
	a, b := d.DaysFromZero(), e.DaysFromZero()
	if a < b {
		return b - a - 1
	} else if a > b {
		return a - b - 1
	}
	return 0
}

// DaysFromZero returns the number of days between the d Date, and the start of the year (i.e. YYYY-01-01)
func (d Date) DaysFromStartOfYear() uint {
	result := d.day

	var daysInYear *[12]uint
	if isLeapYear(d.year) {
		daysInYear = &monthDatesLeap
	} else {
		daysInYear = &monthDates
	}
	for i := 0; i < int(d.month)-1; i++ {
		result += daysInYear[i]
	}
	return result
}

// DaysFromZero returns the number of days between the d Date, and zero (i.e. 0000-00-00)
func (d Date) DaysFromZero() uint {
	result := d.DaysFromStartOfYear()
	if d.year == 0 {
		return result
	}
	result += 365 * (d.year - 1)

	// Compensate for leap years
	result += (d.year / 4) + 1
	return result
}

// Days in each month of Jan, Feb, ..., Dec
var monthDates = [12]uint{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var monthDatesLeap = [12]uint{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
