package main

import (
	"flag"
	"fmt"
	"github.com/Jeadie/DateDiff/dataset"
)

const defaultOutputFilename = "dataset.csv"
const defaultDatasetSize = 100
const defaultPositiveRatio = 0.5

func Dataset(args []string) {
	argFs := flag.NewFlagSet(string(DatasetCmd), flag.ExitOnError)
	param := dataset.DatasetParams{}

	argFs.StringVar(&param.OutputFilename, "out", defaultOutputFilename, "file path to save the output dataset")
	argFs.UintVar(&param.Size, "size", defaultDatasetSize, "size of dataset to construct")
	argFs.Float64Var(&param.PositiveRatio, "labelRatio", defaultPositiveRatio, "proportion of dataset to be valid input dates")
	err := argFs.Parse(args)

	if err != nil {
		fmt.Println(err.Error())
	}

	dataset.CreateDataset(param)
}
