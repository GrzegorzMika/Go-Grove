package devices

import (
	"errors"
	"strconv"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

type DigitalDevice struct {
	GPIOPin gpio.PinIO
}

func (d *DigitalDevice) Type() DeviceType {
	return DeviceTypeDigital
}

func (d *DigitalDevice) Init(pinNumber int) error {
	d.GPIOPin = gpioreg.ByName(strconv.Itoa(pinNumber))
	if d.GPIOPin == nil {
		return errors.New("failed to find GPIO pin")
	}
	return nil
}
