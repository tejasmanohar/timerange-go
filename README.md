# timerange

time range iterators for golang

# usage
main.go
```go
package main

import (
  "fmt"
  "time"

  "github.com/tejasmanohar/timerange-go"
)

func main() {
  start := time.Date(2016, 8, 28, 9, 0, 0, 0, time.UTC)
  end := time.Date(2016, 8, 28, 11, 0, 0, 0, time.UTC)
  iter := timerange.New(start, end, time.Hour)
  for iter.Next() {
    t := iter.Current()
    fmt.Println(t)
  }
}
```

```
$ go run main.go
2016-08-29 09:00:00 +0000 UTC
2016-08-29 10:00:00 +0000 UTC
2016-08-29 11:00:00 +0000 UTC
```
