package main

import (
	"flag"
	"fmt"
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

func testDisplay() {
	var node struct {
		Name   string
		Length int64
	}
	node.Name = "Stephen"
	node.Length = 100
	ch12.Display("node", node)
}
