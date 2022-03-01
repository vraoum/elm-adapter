package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// FuelPressure Define pressure in the fuel system, usually the sensor is in the fuel intake rail in the motor bay
type FuelPressure struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the FuelPressure value, we must apply the following formula: 3 * a
// Where a is the first parameters given
// the range of FuelPressure that this method can return is 0 to 765
func (fp *FuelPressure) Convert(args []string) (string, error) {
	if len(args) == 1 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)

		if decErr1 == nil {
			res := strconv.FormatInt(3*a, 10)
			fp.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 10 Fuel pressure: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (fp *FuelPressure) GetLastValue() string {
	return fp.lastValue
}

// Unit Returns the unit of the pid, for the FuelPressure the unit is kPa as kilopascal
func (fp *FuelPressure) Unit() string {
	return "kPa"
}

// GetPid Returns the pid hex as integer
func (fp *FuelPressure) GetPid() int {
	return 0x0A
}

// GetService Returns the service hex as integer
func (fp *FuelPressure) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (fp *FuelPressure) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", fp.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (fp *FuelPressure) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", fp.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (fp *FuelPressure) GetIsSupported() bool {
	return fp.IsSupported
}
