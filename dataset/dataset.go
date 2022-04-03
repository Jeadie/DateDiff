package dataset

import (
	"encoding/csv"
	"fmt"
	"github.com/Jeadie/DateDiff/diff"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

type DatasetParams struct {
	OutputFilename string
	Size           uint
	PositiveRatio  float64
}

func CreateDataset(param DatasetParams) {
	validCount := uint(float64(param.Size) * param.PositiveRatio)
	invalidCount := param.Size - validCount

	w, err := ConstructWriter(param.OutputFilename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = w.Write([]string{"start", "end", "diff", "valid"})

	out := make(chan []string)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go createLines(validCount, out, ConstructPossibleValidInput, true, wg)
	go createLines(invalidCount, out, ConstructPossibleInvalidInput, false, wg)

	go func(w *csv.Writer) {
		for l := range out {
			w.Write(l)
		}
	}(w)
	wg.Wait()

	w.Flush()
}

func createLines(total uint, out chan []string, inFn func() string, outputValid bool, wg *sync.WaitGroup) {
	defer wg.Done()
	var label string
	if outputValid {
		label = "1"
	} else {
		label = "0"
	}

	for total > 0 {
		s, e := inFn(), inFn()
		v, err := diff.AbsoluteDateDifference(s, e)

		if (err == nil) == outputValid {
			out <- []string{s, e, strconv.Itoa(int(v)), label}
			total--
		}
	}
}

func ConstructWriter(filename string) (*csv.Writer, error) {
	f, e := os.Create(filename)
	if e != nil {
		return &csv.Writer{}, e
	}
	return csv.NewWriter(f), nil
}

func ConstructPossibleInvalidInput() string {
	// Since Uint has 10 base-10 digits, eight least significant digits are entirely random
	v := rand.Uint32()
	return fmt.Sprintf("%04d-%02d-%02d",
		v%5000,
		(v/10000)%25,    // ~50% months will be invalid
		(v/1000000)%100, // of the 50% valid months, ~70% of days invalid
	)
}

func ConstructPossibleValidInput() string {
	// Since Uint has 10 base-10 digits, eight least significant digits are entirely random
	v := rand.Uint32()
	return fmt.Sprintf("%04d-%02d-%02d",
		v%5000,
		(v/1000)%12+1,
		(v/100000)%31+1,
	)
}
