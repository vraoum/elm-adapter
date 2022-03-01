package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ThrottlePosition Define the position of the throttle body inside the engine
type ThrottlePosition struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the ThrottlePosition value, we must apply the following formula: a * 100 / 255
// Where a is the first parameter given
// the range of ThrottlePosition that this method can return is 0 to 100
func (maf *ThrottlePosition) Convert(args []string) (string, error) {
	if len(args) == 1 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)

		if decErr1 == nil {
			res := strconv.FormatInt(a*100/255, 10)
			maf.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 17 Throttle position: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (maf *ThrottlePosition) GetLastValue() string {
	return maf.lastValue
}

// Unit Returns the unit of the pid, for the MassAirFlowSensor the unit is percentage
func (maf *ThrottlePosition) Unit() string {
	return "%"
}

// GetPid Returns the pid hex as integer
func (maf *ThrottlePosition) GetPid() int {
	return 0x11
}

// GetService Returns the service hex as integer
func (maf *ThrottlePosition) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (maf *ThrottlePosition) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", maf.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (maf *ThrottlePosition) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", maf.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (maf *ThrottlePosition) GetIsSupported() bool {
	return maf.IsSupported
}
