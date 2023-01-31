package reciever

import (
	"exercise/models"
	"fmt"
	"sync"
)

type Reciever struct {
	satellites *[]models.Satellite
}

func (r Reciever) Initialize(satellites ...*models.Satellite) Reciever {
	for _, satellite := range satellites {
		r.satellites = append(r.satellites, satellite)
	}
	return r
}

func (r Reciever) RecieveAllMessages(channels []chan string) {
	wp := sync.WaitGroup{}

	for i, channel := range channels {
		wp.Add(1)
		go func(chan2 chan string, index int) {
			defer wp.Done()

			fmt.Println("Address -> ", &r.satellites)

			satellite := *r.satellites
			satellite[index-1].SetMessage(chan2)

		}(channel, i)

		wp.Wait()

		fmt.Println("All messages has been received")
	}
}

func (r Reciever) PrintMessages() {
	for _, satellite := range *r.satellites {
		fmt.Printf("Messages coming from %s -> /n/t %s", satellite.Id, satellite.GetMessages())
	}
}
