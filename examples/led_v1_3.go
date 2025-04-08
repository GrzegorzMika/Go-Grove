package main

import (
	"fmt"
	"time"

	"github.com/GrzegorzMika/Go-Grove/grove"
	"github.com/gSpera/morse"
)

var MESSAGE = "Hello, World!"

func morseToLED(morseCode string, led grove.LEDv1_3) {
	for _, char := range morseCode {
		switch char {
		case '.':
			led.TurnOn()
			time.Sleep(100 * time.Millisecond)
			led.TurnOff()
			time.Sleep(100 * time.Millisecond)
		case '-':
			led.TurnOn()
			time.Sleep(300 * time.Millisecond)
			led.TurnOff()
			time.Sleep(100 * time.Millisecond)
		case ' ':
			time.Sleep(700 * time.Millisecond)
		}
	}
}

func main() {
	led := grove.LEDv1_3{}
	err := led.Init(22)
	if err != nil {
		fmt.Println(err)
	}

	// Convert to morse
	textInMorse := morse.ToMorse(MESSAGE)
	morseToLED(textInMorse, led)
}
