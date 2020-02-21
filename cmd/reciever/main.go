package main

import (
	"fmt"
	"os"

	"github.com/ripx80/signal"
)

type SniffOptions struct {
	gpioPin uint
}

func (o *SniffOptions) Run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	receiver := signal.NewReceiver(o.gpioPin)
	defer receiver.Close()

	for {
		select {
		case res := <-receiver.Receive():
			fmt.Printf("received code=%d pulseLength=%d bitLength=%d protocol=%d\n",
				res.Code, res.PulseLength, res.DefaultBitLength, res.Protocol)
		case <-interrupt:
			fmt.Println("received interrupt")
			return
		}
	}
}

func main() {
	options := &SniffOptions{
		gpioPin: 17,
	}
	options.Run()
}
