package main

import (
	"github.com/tpisntgod/Service_Calculation/Hw4_Week12/cloudgo-xorm/service"
)

func main() {
	n := service.NewServer()
	n.Run(":8080")
}
