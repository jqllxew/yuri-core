package start

import (
	"fmt"

	"github.com/jqllxew/YuriCore/conf"
	"github.com/jqllxew/YuriCore/main_service/client"
	"github.com/jqllxew/YuriCore/main_service/infrastructure/room_service"
)

func initRoomService() {
	client.InitRoomClient(room_service.NewRoomServiceImpl(fmt.Sprintf("%s:%d", conf.Config.RoomAdress, conf.Config.RoomPort)))
}
