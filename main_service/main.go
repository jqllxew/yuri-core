package main

import (
	"fmt"
	"github.com/jqllxew/YuriCore/main_service/start"
	"github.com/jqllxew/YuriCore/utils"
	"os"
	"os/signal"
	"syscall"
)

var (
	SERVERVERSION = "v1.0"
)

func main() {
	fmt.Println("YuriCore Server", SERVERVERSION)
	fmt.Println("Initializing process ...")

	ExePath, err := utils.GetExePath()
	if err != nil {
		panic(err)
	}
	start.Init(ExePath)

	go initTCP()
	go initUDP()

	ch := make(chan os.Signal)
	defer close(ch)
	signal.Notify(ch, syscall.SIGINT)
	_ = <-ch

	fmt.Println("Press CTRL+C again to close server")

	signal.Notify(ch, syscall.SIGINT)
	_ = <-ch
}
