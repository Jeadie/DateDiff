# DateDiff
DateDiff does one thing, and one thing well: calculate the absolute differences, in days, between two dates of the exact format `YYYY-MM-DD`.

## Usage
There are three ways to use datediff: a library, via the command line or as a server (either self-configured, or as a docker container).

### Library
```go
import (
    "github.com/Jeadie/DateDiff/diff"
)

days, err := diff.AbsoluteDateDifference("2020-02-20", "2002-02-20")
if err == nil {
	fmt.Printf("%d days\n", days)
}
```

### Command Line
After installing (see [install binary]()):
```shell
 datediff cmp 2020-02-20 2002-02-20
```

### Server
For running from binary, after installing (see [install binary]()):
```shell
datediff server -p=8001
```

**(WIP)**: Or the server can be run as a Docker container:
```shell
docker run ...
```

## Installation
> WORK IN PROGRESS