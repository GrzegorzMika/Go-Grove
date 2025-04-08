package devices

import "strings"

const DuplicateGPIORegistrationError = `can't register pin "GPIO0" twice; already registered as "GPIO0"`

func filterError(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), DuplicateGPIORegistrationError) {
		return nil
	}
	return err
}
