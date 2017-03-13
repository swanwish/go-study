package main

import (
	"fmt"
	"strings"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-common/utils"
)

func main() {
	ips, err := utils.GetLocalIPAddrs()
	if err != nil {
		logs.Errorf("Failed to local ip address, the error is %v", err)
		return
	}
	fmt.Printf("The local ips are: %s\n", strings.Join(ips, ", "))
}
