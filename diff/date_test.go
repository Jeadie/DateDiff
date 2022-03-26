package diff

import "testing"

type DateTest struct {
	start                      string
	end                        string
	expectedAbsoluteDifference uint // Not checked if isErrorExpected is true
	isErrorExpected            bool
	failureMessage             string // Empty if isErrorExpected is false
}

const IrrelevantAbsoluteDifferenceValue = 0

var dateTests = []DateTest{
	{
		"",
		"",
		IrrelevantAbsoluteDifferenceValue,
		true,
		"empty date strings are invalid",
	},
	{
		"20000-01-01",
		"20000-01-01",
		IrrelevantAbsoluteDifferenceValue,
		true,
		"invalid year",
	},
	{
		"2000-13-01",
		"2000-13-01",
		IrrelevantAbsoluteDifferenceValue,
		true,
		"invalid month",
	},
	{
		"2000-01-40",
		"2000-01-40",
		IrrelevantAbsoluteDifferenceValue,
		true,
		"invalid day",
	},
	{
		"01-01-2020",
		"01-01-2020",
		IrrelevantAbsoluteDifferenceValue,
		true,
		"invalid date structure. Expected structure YYYY-MM-DD",
	},
	{
		"2022-02-29",
		"2022-02-29",
		IrrelevantAbsoluteDifferenceValue,
		true,
		"invalid date, There is no 29th February in 2022",
	},
	{
		"2022/01/01",
		"2022/01/01",
		IrrelevantAbsoluteDifferenceValue,
		true,
		"invalid date structure. Expected structure YYYY-MM-DD",
	},
	{
		"2012-01-10",
		"2012-01-11",
		0,
		false,
		"",
	},
	{
		"2012-01-01",
		"2012-01-10",
		8,
		false,
		"",
	},
	{
		"1801-06-13",
		"1801-11-11",
		150,
		false,
		"",
	},
	{
		"2021-12-01",
		"2017-12-14",
		1447,
		false,
		"",
	},
	{
		"2017-12-14",
		"2021-12-01",
		1447,
		false,
		"",
	},
}

func TestDate(t *testing.T) {
	for _, tt := range dateTests {
		diff, err := AbsoluteDateDifference(tt.start, tt.end)

		// Check for error
		if !tt.isErrorExpected && err != nil {
			t.Errorf("AbsoluteDateDifference(%s, %s) did not expect an error, but error returned: %s", tt.start, tt.end, err)
		} else if tt.isErrorExpected && err == nil {
			t.Errorf("AbsoluteDateDifference(%s, %s) did not return an error, but an error was expected because %s", tt.start, tt.end, tt.failureMessage)
		}

		// Check for correct difference
		if diff != tt.expectedAbsoluteDifference {
			t.Errorf("%s <-> %s has an absolute difference of %d, but returned %d", tt.start, tt.end, tt.expectedAbsoluteDifference, diff)
		}
	}
}
