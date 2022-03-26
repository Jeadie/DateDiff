package main

import (
	"fmt"
	"github.com/Jeadie/DateDiff/diff"
	"os"
	"strconv"
)

func Cmp(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "`cmp` requires two input arguments. %d given\n", len(args))
		return
	}
	start, end := args[0], args[1]
	result, err := diff.AbsoluteDateDifference(start, end)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	} else {
		fmt.Fprintf(os.Stdout, strconv.Itoa(int(result)))
	}
}
