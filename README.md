# DateDiff
DateDiff does one thing, and one thing well: calculate the absolute difference, in days, between two dates of the exact format `YYYY-MM-DD`.

## Usage
There are three ways to use datediff:
- as a Go library
- via the command line
- or run it as a server

### Library
Datediff has the go module `github.com/Jeadie/DateDiff/diff`. Simply import and use `diff.AbsoluteDateDifference`.
```go
import (
    "github.com/Jeadie/DateDiff/diff"
)

days, err := diff.AbsoluteDateDifference("2020-02-20", "2002-02-20"); if err == nil {
	fmt.Printf("%d days\n", days)
}
```

### Command Line
Datediff can also run via the terminal. After installing or building from source. Simply run
```shell
 datediff cmp 2020-02-20 2002-02-21
```

### Server
Datediff can operate as a server. Either run the binary in server mode:
```shell
datediff server -p=8001
```

or the server can be run as a Docker container:
```shell
docker run -e PORT=8001 -t jeadi/datediff:latest
```

## Installation
Simple, use the Makefile
```shell
make build
```

Binaries are also available for each release and are built against multiple common OS's and architectures. Find them [here](https://github.com/Jeadie/DateDiff/releases/). To download the latest binary for the local OS/arch:
```shell 
make download
 
# Binary can then be used as above, e.g. 
./datediff cmp 2020-02-20 2002-02-20
```