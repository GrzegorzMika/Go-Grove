package grove

import (
	"github.com/GrzegorzMika/Go-Grove/devices"
	"periph.io/x/conn/v3/gpio"
)

type LEDv1_3 struct {
	*devices.DigitalDevice
}

func (d *LEDv1_3) TurnOn() error {
	return d.GPIOPin.Out(gpio.High)
}
func (d *LEDv1_3) TurnOff() error {
	return d.GPIOPin.Out(gpio.Low)
}
