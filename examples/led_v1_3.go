package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/GrzegorzMika/Go-Grove/grove"
	"github.com/gSpera/morse"
	"periph.io/x/host/v3"
)

var MESSAGE = "Hello, World!"

func morseToLED(morseCode string, led *grove.LEDv1_3) {
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
	state, err := host.Init()
	if err != nil {
		fmt.Println("Error initializing host:", err)
	}
	if len(state.Failed) > 0 {
		errs := make([]error, len(state.Failed))
		for i, f := range state.Failed {
			errs[i] = f.Err
		}
		fmt.Println("Errors initializing host:", errors.Join(errs...))
	}

	led, err := grove.NewLEDv1_3(22)
	if err != nil {
		fmt.Println("Error initializing LED:", err)
	}

	// Convert to morse
	textInMorse := morse.ToMorse(MESSAGE)
	morseToLED(textInMorse, led)
}
