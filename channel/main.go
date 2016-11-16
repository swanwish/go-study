package main

import (
	"time"

	"github.com/swanwish/go-common/logs"
)

var channelContainer = make(chan int64, 10)

func init() {
	//go consumeData()
}

func main() {
	go produceData()
	consumeData()
}

func produceData() {
	var index int64
	for index = 1; index < 100; index++ {
		channelContainer <- index
		logs.Debugf("Item %d added to channel", index)
	}
}

func consumeData() {
	for {
		select {
		case item := <-channelContainer:
			logs.Debugf("consume item %d", item)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
