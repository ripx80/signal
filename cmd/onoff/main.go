package main

import (
	"fmt"
	"time"

	"github.com/ripx80/gpio"
)

/*
transmit sends a on and off signal. Change the codes
*/

type transmitOptions struct {
	PulseLength uint
	GpioPin     uint
	Protocol    int
	BitLength   int
}

func pb(code uint64, len int) {
	fmt.Printf(fmt.Sprintf("sending {%%d}: %%0%db \n", len),
		len,
		code,
	)
}

func main() {

	options := &transmitOptions{
		PulseLength: 330,
		GpioPin:     gpio.DefaultTransmitPin,
		Protocol:    gpio.DefaultProtocol,
		BitLength:   gpio.DefaultBitLength,
	}

	var on, off uint64
	on = 0b000000000001010100010101  //on 5397
	off = 0b000000000001010100010100 // off 5396

	pb(on, 24)

	t := gpio.NewTransmitter(options.GpioPin)

	err := t.Transmit(on, options.Protocol, options.PulseLength, options.BitLength)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Wait()
	time.Sleep(2 * time.Second)

	pb(off, 24)
	err = t.Transmit(off, options.Protocol, options.PulseLength, options.BitLength)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Wait()
}
