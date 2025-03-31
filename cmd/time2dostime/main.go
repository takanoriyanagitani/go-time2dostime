package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	td "github.com/takanoriyanagitani/go-time2dostime"
	. "github.com/takanoriyanagitani/go-time2dostime/util"
)

var args IO[[]string] = Of(os.Args[1:])

var timeFromArgRfc3339 IO[string] = Bind(
	args,
	Lift(func(s []string) (string, error) {
		switch len(s) {
		case 1:
			return s[0], nil
		default:
			return "", fmt.Errorf("unexpected arg len: %v", len(s))
		}
	}),
)

type StringToTime func(string) (time.Time, error)

func LayoutToTimeParser(layout string) StringToTime {
	return func(s string) (time.Time, error) {
		return time.Parse(layout, s)
	}
}

var str2timeRfc3339 StringToTime = LayoutToTimeParser(time.RFC3339)

var timeFromArg IO[time.Time] = Bind(
	timeFromArgRfc3339,
	Lift(str2timeRfc3339),
)

var timeOrNow IO[time.Time] = timeFromArg.Or(Of(time.Now()))

var basicTime IO[td.BasicDateTime] = Bind(
	timeOrNow,
	Lift(func(t time.Time) (td.BasicDateTime, error) {
		return td.TimeToBasic(t), nil
	}),
)

var dostimeUnsigned IO[uint32] = Bind(
	basicTime,
	Lift(func(b td.BasicDateTime) (uint32, error) { return b.DosTime(), nil }),
)

var dostime2stdout func(uint32) IO[Void] = func(d uint32) IO[Void] {
	return func(_ context.Context) (Void, error) {
		_, e := fmt.Println(d)
		return Empty, e
	}
}

var argnow2dostime2stdout IO[Void] = Bind(
	dostimeUnsigned,
	dostime2stdout,
)

func main() {
	_, e := argnow2dostime2stdout(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
