package grove

import (
	"github.com/GrzegorzMika/Go-Grove/devices"
	"periph.io/x/conn/v3/gpio"
)

type LEDv1_3 struct {
	*devices.DigitalDevice
}

func NewLEDv1_3(pinNumber int) (*LEDv1_3, error) {
	led := &LEDv1_3{
		DigitalDevice: &devices.DigitalDevice{
			Device: &devices.Device{},
		},
	}
	err := led.Init(pinNumber)
	if err != nil {
		return nil, err
	}
	return led, nil
}

func (d *LEDv1_3) TurnOn() error {
	return d.GPIOPin.Out(gpio.High)
}
func (d *LEDv1_3) TurnOff() error {
	return d.GPIOPin.Out(gpio.Low)
}
