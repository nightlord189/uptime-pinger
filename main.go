package main

import (
	"fmt"
	"github.com/nightlord189/uptime-pinger/pinger"
)

func main() {
	fmt.Println("start")
	pingerInstance := pinger.Pinger{}
	result := pingerInstance.CheckUrl("https://google.com")
	fmt.Println(result)
}
