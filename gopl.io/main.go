package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-study/gopl.io/ch12"
)

var op string

func main() {
	flag.StringVar(&op, "op", "display", "The operation to execute")
	flag.Parse()
	switch op {
	case "display":
		testDisplay()
	case "format":
		testFormat()
	default:
		logs.Errorf("Unknown operation")
	}
}

func testFormat() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(ch12.Any(x))
	fmt.Println(ch12.Any(d))
	fmt.Println(ch12.Any([]int64{x}))
	fmt.Println(ch12.Any([]time.Duration{d}))
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func testDisplay() {
	strangeLove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		}}
	ch12.Display("strangeLove", strangeLove)

	ch12.Display("os.Stderr", os.Stderr)

	ch12.Display("rV", reflect.ValueOf(os.Stderr))
}
