package main

import (
	"fmt"
	"github.com/jqllxew/YuriCore/utils"

	"github.com/jqllxew/YuriCore/conf"
	"github.com/jqllxew/YuriCore/user_service/start"
)

func main() {
	ExePath, err := utils.GetExePath()
	if err != nil {
		panic(err)
	}
	start.Init(ExePath)
	initServer(ExePath, fmt.Sprintf("%s:%d", conf.Config.UserAdress, conf.Config.UserPort))
}
