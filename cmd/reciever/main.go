package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ripx80/gpio"
)

type SniffOptions struct {
	GpioPin uint
}

func (o *SniffOptions) Run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	receiver := gpio.NewReceiver(o.GpioPin)
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
		GpioPin: 17,
	}
	options.Run()
}
