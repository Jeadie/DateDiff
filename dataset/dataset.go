package dataset

import "fmt"

type DatasetParams struct {
	OutputFilename string
	Size           uint
	PositiveRatio  float64
}

func CreateDataset(param DatasetParams) {
	fmt.Print(param)
}
