package devices

import (
	"errors"
	"strconv"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type DigitalDevice struct {
	GPIOPin gpio.PinIO
}

func (d *DigitalDevice) Type() DeviceType {
	return DeviceTypeDigital
}

func (d *DigitalDevice) Init(pinNumber int) error {
	state, err := host.Init()
	err = filterError(err)
	if err != nil {
		return err
	}
	if len(state.Failed) > 0 {
		errs := make([]error, len(state.Failed))
		for i, f := range state.Failed {
			errs[i] = f.Err
		}
		return errors.Join(errs...)
	}
	d.GPIOPin = gpioreg.ByName(strconv.Itoa(pinNumber))
	if d.GPIOPin == nil {
		return errors.New("failed to find GPIO pin")
	}
	return nil
}
