package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	strToday := time.Now().Format("20060102")
	fmt.Println(now, strToday)
}
