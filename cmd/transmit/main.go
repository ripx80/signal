package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ripx80/signal"
)

/*
transmit sends a on and off signal. Change the codes
GOARM=6 GOARCH=arm GOOS=linux go build -o transmit main.go
*/

type transmitOptions struct {
	PulseLength uint
	signalPin   uint
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

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <code>\n", os.Args[0])
		return
	}
	var code uint64
	i, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	code = uint64(i)

	options := &transmitOptions{
		PulseLength: 330,
		gpioPin:     signal.DefaultTransmitPin,
		Protocol:    signal.DefaultProtocol,
		BitLength:   signal.DefaultBitLength,
	}

	pb(code, 24)

	t := signal.NewTransmitter(options.gpio.Pin)

	err = t.Transmit(code, options.Protocol, options.PulseLength, options.BitLength)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Wait()
}
