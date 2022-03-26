package DateDiff

import "testing"

type UintParseTest struct {
	x               string
	y               uint
	isErrorExpected bool
	invalidMessage  string
}

var uintParseTests = []UintParseTest{
	{
		"",
		0,
		true,
		"empty input is invalid",
	},
	{
		"abc1",
		0,
		true,
		"contains invalid digits",
	},
	{
		"-10",
		0,
		true,
		"is a signed integer",
	},
	{
		"10.0",
		0,
		true,
		"is a decimal number",
	},
	{
		"1",
		1,
		false,
		"",
	},
	{
		"12",
		12,
		false,
		"",
	},
	{
		"123",
		123,
		false,
		"",
	},
}

func TestUintParse(t *testing.T) {
	for _, tt := range uintParseTests {
		actual, err := UintParse(tt.x)

		if tt.isErrorExpected && err == nil {
			// Didn't error
			t.Errorf("UintParse(%s) expected, but did not yield an error because %s", tt.x, tt.invalidMessage)

		} else if !tt.isErrorExpected && err != nil {
			// Errored for some reason
			t.Errorf("UintParse(%s) produced an unexpected error %s", tt.x, err)
		}

		if tt.y != actual {
			t.Errorf("UintParse(%s) was parsed incorrectly. Expected %d, found %d", tt.x, tt.y, actual)
		}

	}
}
