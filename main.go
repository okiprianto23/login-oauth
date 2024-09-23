package main

import (
	"github.com/okiprianto23/login-oauth/config"
	"github.com/okiprianto23/login-oauth/serverconfig"
	"github.com/okiprianto23/zerolog/log"
)

func main() {
	appConfig := config.AppConfig
	serverAttribute := serverconfig.NewServerAttribute(appConfig)
	err := serverAttribute.Init()
	defer func() {
		//TODO close connection
	}()
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}
}
