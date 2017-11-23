package main

import "github.com/tpisntgod/Service_Calculation/cloudgo-io/service"

func main() {
	n := service.NewServer()
	n.Run(":8080")
}
