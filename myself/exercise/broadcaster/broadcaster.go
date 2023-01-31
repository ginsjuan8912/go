package broadcaster

import (
	"bufio"
	"context"
	"errors"
	"exercise/models"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Broadcaster struct {
	X         float32
	Y         float32
	receivers []chan string
}

func (b Broadcaster) Initialize(satellites ...models.Satellite) Broadcaster {
	for i := range satellites {
		channel := make(chan string, i)
		b.receivers = append(b.receivers, channel)
	}

	return b
}
func (b Broadcaster) BroadcastMessages() error {

	ctx, cancelOperation := context.WithCancel(context.Background())

	if len(b.receivers) < 0 {
		return errors.New("Cannot start broadcasting without any reciever!")
	}

	fmt.Println("Broadcaster: Sending messages...")
	go func() {
		fmt.Println("Press ENTER to cancel the broadcast...")
		fmt.Scanf("/n")
		cancelOperation()
	}()

	broadcastDone := make(chan int8)

	go func(context context.Context) {
		defer close(broadcastDone)

		file, err := os.Open("message.txt")

		if err != nil {
			panic("cannot read the requested file")
		}

		defer file.Close()

		reader := bufio.NewReader(file)

		maxRecievers := len(b.receivers)
		minRecievers := 1

		for {
			line, _, err := reader.ReadLine()

			if err == io.EOF {
				break
			}

			for _, s := range strings.Split(string(line), ",") {

				if s == "0 " || s == "\n" {
					continue
				}

				b.receivers[generateNumber(minRecievers, maxRecievers)-1] <- s
			}

		}

		for _, receiver := range b.receivers {
			close(receiver)
		}

		broadcastDone <- 1

	}(ctx)

	<-broadcastDone

	return nil
}

func generateNumber(min int, max int) int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixMilli())
	return rand.Intn(max-min+1) + min
}
