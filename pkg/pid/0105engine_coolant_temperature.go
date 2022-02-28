package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// EngineCoolantTemperature Define the temperature of the coolant used to cold the motor
type EngineCoolantTemperature struct {
	lastValue string // lastValue Last value received
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the EngineCoolantTemperature value, we must apply the following formula: a - 40
// Where a is the first parameters given
// the range of EngineCoolantTemperature that this method can return is -40 to 215
func (ect *EngineCoolantTemperature) Convert(args []string) (string, error) {
	if len(args) == 1 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)

		if decErr1 == nil {
			res := strconv.FormatInt(a-40, 10)
			ect.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (ect *EngineCoolantTemperature) GetLastValue() string {
	return ect.lastValue
}

// Unit Returns the unit of the pid, for the EngineCoolantTemperature the unit is °C as Celsius degree
func (ect *EngineCoolantTemperature) Unit() string {
	return "°C"
}

// GetPid Returns the pid hex as integer
func (ect *EngineCoolantTemperature) GetPid() int {
	return 0x05
}

// GetService Returns the service hex as integer
func (ect *EngineCoolantTemperature) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (ect *EngineCoolantTemperature) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", ect.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (ect *EngineCoolantTemperature) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", ect.GetService()))
}
