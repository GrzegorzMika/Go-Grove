package devices

type DigitalDevice struct {
	Device
}

func (d *DigitalDevice) Type() DeviceType {
	return DeviceTypeDigital
}
