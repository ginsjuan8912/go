package main

import (
	broadcaster2 "exercise/broadcaster"
	"exercise/models"
	"exercise/reciever"
)

func main() {

	kenobi := models.Satellite{Id: "Kenobi", X: -500, Y: -200}
	skywalker := models.Satellite{Id: "Skywalker", X: 100, Y: -100}
	sato := models.Satellite{Id: "Sato", X: 500, Y: 100}

	broadcast := broadcaster2.Broadcaster{}
	recieverHandler := reciever.Reciever{}
	recieverHandler = reciever.Reciever{}.Initialize()
	broadcast = broadcast.Initialize(kenobi, skywalker, sato)
	err := broadcast.BroadcastMessages()

	if err != nil {
		return
	}

}
