package dataset

import (
	"encoding/csv"
	"fmt"
	"github.com/Jeadie/DateDiff/diff"
	"math/rand"
	"os"
	"strconv"
)

type DatasetParams struct {
	OutputFilename string
	Size           uint
	PositiveRatio  float64
}

func CreateDataset(param DatasetParams) {
	validCount := uint(float64(param.Size) * param.PositiveRatio)
	invalidCount := param.Size - validCount
	fmt.Println(validCount, invalidCount)

	w, err := ConstructWriter(param.OutputFilename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = w.Write([]string{"start", "end", "diff", "valid"})

	for validCount > 0 || invalidCount > 0 {
		s, e := ConstructInput(), ConstructInput()
		v, err := diff.AbsoluteDateDifference(s, e)
		if err == nil {
			if validCount > 0 {
				_ = w.Write([]string{s, e, strconv.Itoa(int(v)), "1"})
				validCount--
			}
		} else {
			if invalidCount > 0 {
				// Invalid difference, use placeholder difference value
				_ = w.Write([]string{s, e, "0", "0"})
				invalidCount--
			}
		}
	}
	w.Flush()
}

func ConstructWriter(filename string) (*csv.Writer, error) {
	f, e := os.Create(filename)
	if e != nil {
		return &csv.Writer{}, e
	}
	return csv.NewWriter(f), nil
}

func ConstructInput() string {
	// Since Uint has 10 base-10 digits, eight least significant digits are entirely random
	v := rand.Uint32()
	return fmt.Sprintf("%04d-%02d-%02d",
		v%1000,
		(v/1000)%100,
		(v/100000)%100,
	)
}
