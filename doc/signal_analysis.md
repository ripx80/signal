# Signal Analysis

Sample rate set to 250000 S/s.
Tuner gain set to Auto.
Tuned to 433.920MHz.

## Original Unit

IT-3500L Schalter B, Systemcode
original signal: [00] {65} b3 53 54 d4 b3 55 4b 2b 00
retries: 5
signal_start = 551242, signal_end = 669413, signal_len = 118171, pulses_found = 330
Pulse length 55

// for intertechno mySwitch.setPulseLength(320);

## 433 rc-switch

original signal: {33} 00 15 15 41 00 : 00000000 00010101 00010101 01000001 0
retries: 10
Pulse coding: Short pulse length 127 - Long pulse length 301
signal_start = 39017331, signal_end = 39198728, signal_len = 181397, pulses_found = 330

## GO gpio

original signal: {25} 15 15 41 00 : 00010101 00010101 01000001 0
retries: 10
signal_start = 92630376, signal_end = 92724223, signal_len = 93847, pulses_found = 250
Pulse coding: Short pulse length 78 - Long pulse length 174

## Manual

433:    00000000 00010101 00010101 01000001 0
decode  00000000 00010101 00010101 01010001 0
        00000000 00010101 00010101 01000001 0

gpio  0000000000010101 00010101 01000001
decode        10010101 00010101 01000001 0

## From Go

```go

//on := []byte{0xb3, 0x53, 0x54, 0xd4, 0xb3, 0x55, 0x4b, 0x2a, 0x80} // original from hackrf
//on := []byte{0x00, 0x15, 0x15, 0x41}
/*
            ?binary: 10110011 1010011 1010100 11010100 10110011 1010101 1001011 101010 10000000

            cybercf 10110011 01010011 01010100 11010100 10110011 01010101 01001011 00101011 00000000

            hex:	b3 53 54 d4 b3 55 4b 2b 00
                    b3 53 d4 d4 b3 55 4b 2b 00
                    b3 53 54 d4 b3 55 4b 2b 00

            443cmd
            hex: 	00 15 15 41 00
            binary:	00000000 00010101 00010101 01000001 0
            binary  00000000 00010101 00010101 01000001 // im c code


    1011001101010011010101001101010010110011010101010100101100101011
    1011001101010011010101001101010010110011010101010100101100101011
*/

//off := []byte{0xb3, 0x53, 0x54, 0xd4, 0xb3, 0x55, 0x4b, 0xaa, 0x80}
//off := []byte{0x15, 0x14, 0x54}
/*
    binary: 10110011 1010011 1010100 11010100 10110011 1010101 1001011 10101010 10000000
    cybercf	10110011 01010011 01010100 11010100 10110011 01010101 01001010 10101010 10000000
    uint64:	10110011 01010011 01010100 11010100 10110011 01010101 01001011 10101010
    hex: 	b3 53 54 d4 b3 55 4a ab 00

    443cmd
    hex: 	00 15 14 54 00
    binary: 00000000 00010101 00010100 01010100 0
    binary: 00000000 00010101 00010100 01010100 // im c code
*/

//codeOn := binary.BigEndian.Uint64(on)
//codeOff := binary.BigEndian.Uint64(off)

//fmt.Printf("Len: %d\tCodeOn:\t %0b\n", bits.Len64(codeOn), codeOn)
//fmt.Printf("Len: %d\tCodeOff: %0b\n", bits.Len64(codeOff), codeOff)
//code := codeOn
```

## External

https://www.tinkerforge.com/de/shop/bricklets/remote-switch-v2-bricklet.html

try this out!!!
https://wiki.fhem.de/wiki/ITLR-3500_/_ITLR-3500T
https://wiki.kainhofer.com/hardware/vaillantvrt340f
https://github.com/miek/inspectrum
https://wiki.fhem.de/wiki/Intertechno_Code_Berechnung#Selbstlernende_Intertechno_Funksteckdosen_.28z.B._ITR-1500.29