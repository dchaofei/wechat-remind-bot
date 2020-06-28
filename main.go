package main

import (
	"github.com/dchaofei/wechat-remind-bot/models"
	"github.com/dchaofei/wechat-remind-bot/startup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	startup.SetupVars()
	models.Setup()
	var quitSig = make(chan os.Signal)
	signal.Notify(quitSig, os.Interrupt, syscall.SIGTERM)
	select {
	case <-quitSig:
		log.Fatal("exit.by.signal")
	}
}
