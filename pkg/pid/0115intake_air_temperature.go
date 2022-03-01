package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// IntakeAirTemperature Define the temperature of the air in the intake
type IntakeAirTemperature struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the IntakeAirTemperature value, we must apply the following formula: a - 40
// Where a is the first parameters given
// the range of IntakeAirTemperature that this method can return is -40 to 215
func (iat *IntakeAirTemperature) Convert(args []string) (string, error) {
	if len(args) == 1 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)

		if decErr1 == nil {
			res := strconv.FormatInt(a-40, 10)
			iat.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 05 Engine coolant temperature: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (iat *IntakeAirTemperature) GetLastValue() string {
	return iat.lastValue
}

// Unit Returns the unit of the pid, for the IntakeAirTemperature the unit is °C as Celsius degree
func (iat *IntakeAirTemperature) Unit() string {
	return "°C"
}

// GetPid Returns the pid hex as integer
func (iat *IntakeAirTemperature) GetPid() int {
	return 0x0F
}

// GetService Returns the service hex as integer
func (iat *IntakeAirTemperature) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (iat *IntakeAirTemperature) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", iat.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (iat *IntakeAirTemperature) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", iat.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (iat *IntakeAirTemperature) GetIsSupported() bool {
	return iat.IsSupported
}
