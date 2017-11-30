package main

import (
	"github.com/tpisntgod/Service_Calculation/cloudgo-io/model"
	"github.com/tpisntgod/Service_Calculation/cloudgo-io/service"
)

func main() {
	n := service.NewServer()
	defer model.Db.Close()
	n.Run(":8080")
}
