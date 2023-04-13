package main

import (
	"fmt"
	"github.com/jqllxew/YuriCore/conf"
	"github.com/jqllxew/YuriCore/room_service/start"
	"github.com/jqllxew/YuriCore/utils"
)

func main() {
	ExePath, err := utils.GetExePath()
	if err != nil {
		panic(err)
	}
	start.Init(ExePath)
	initServer(ExePath, fmt.Sprintf("%s:%d", conf.Config.RoomAdress, conf.Config.RoomPort))
}
