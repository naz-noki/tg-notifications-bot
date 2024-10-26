package main

import (
	httpServer "25ugleyBot/internal/delivery/http"
	"25ugleyBot/pkg/configParser"
	"25ugleyBot/pkg/logger"
	"log"
)

func main() {
	errConfigParser := configParser.New("config", "./configs/", "yaml")
	if errConfigParser != nil {
		log.Fatalln(errConfigParser)
	}

	logFile, errLogger := logger.New(configParser.Cfg.App.Mode, configParser.Cfg.App.LogFilePath)
	if errLogger != nil {
		log.Fatalln(errLogger)
	}
	defer logFile.Close()

	errStart := httpServer.Start(configParser.Cfg.Server.Host, configParser.Cfg.Bot.Token, configParser.Cfg.Bot.ChatId, configParser.Cfg.Server.Port)
	if errStart != nil {
		log.Fatalln(errStart)
	}
}
