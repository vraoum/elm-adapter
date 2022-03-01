package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// MassAirFlowSensor Define the amount air the pass through the sensor per second
// The mass air flow sensor is usually situated just before the air filter
type MassAirFlowSensor struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the MassAirFlowSensor value, we must apply the following formula: (256 * a + b)/100
// Where a and b are the first and second parameters given
// the range of MassAirFlowSensor that this method can return is 0 to 655.35
func (maf *MassAirFlowSensor) Convert(args []string) (string, error) {
	if len(args) == 2 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)
		b, decErr2 := strconv.ParseInt(args[1], 16, 64)

		if decErr1 == nil && decErr2 == nil {
			res := strconv.FormatInt((256*a+b)/100, 10)
			maf.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 16 Mass air flow sensor: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (maf *MassAirFlowSensor) GetLastValue() string {
	return maf.lastValue
}

// Unit Returns the unit of the pid, for the MassAirFlowSensor the unit is g/sec as grams per second
func (maf *MassAirFlowSensor) Unit() string {
	return "g/sec"
}

// GetPid Returns the pid hex as integer
func (maf *MassAirFlowSensor) GetPid() int {
	return 0x10
}

// GetService Returns the service hex as integer
func (maf *MassAirFlowSensor) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (maf *MassAirFlowSensor) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", maf.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (maf *MassAirFlowSensor) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", maf.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (maf *MassAirFlowSensor) GetIsSupported() bool {
	return maf.IsSupported
}
