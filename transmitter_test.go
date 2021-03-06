package signal_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ripx80/signal"
	"github.com/stretchr/testify/assert"
)

type testPin struct {
	sequence bytes.Buffer
	closed   bool
}

func newTestPin() *testPin {
	return &testPin{}
}

func (p *testPin) High() error {
	p.sequence.WriteRune('1')
	return nil
}

func (p *testPin) Low() error {
	p.sequence.WriteRune('0')

	return nil
}

func (p *testPin) Close() {
	p.closed = true
}

func TestTransmitterTransmit(t *testing.T) {
	signal.TransmitRetries = 1
	pin := newTestPin()

	transmitter := signal.NewNativeTransmitter(pin)
	defer transmitter.Close()

	err := transmitter.Transmit(0x1, 1, 190, 24)

	assert.Nil(t, err)

	transmitter.Wait()

	assert.Equal(t, "10101010101010101010101010101010101010101010101010", pin.sequence.String())
}

func TestTransmitInvalidProtocol(t *testing.T) {
	pin := newTestPin()

	transmitter := signal.NewNativeTransmitter(pin)
	defer transmitter.Close()

	err := transmitter.Transmit(0x1, 999, 190, 24)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("Protocol 999 does not exist"), err)
	}
}

func TestTransmitterClose(t *testing.T) {
	pin := newTestPin()

	transmitter := signal.NewNativeTransmitter(pin)
	assert.Nil(t, transmitter.Close())

	assert.True(t, pin.closed)
}

func TestNullTransmitInvalidProtocol(t *testing.T) {
	transmitter := signal.NewNullTransmitter()
	defer transmitter.Close()

	err := transmitter.Transmit(0x1, 999, 190, 24)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("Protocol 999 does not exist"), err)
	}
}
