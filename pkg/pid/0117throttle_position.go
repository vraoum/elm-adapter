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
func (tp *ThrottlePosition) Convert(args []string) (string, error) {
	if len(args) == 1 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)

		if decErr1 == nil {
			res := strconv.FormatInt(a*100/255, 10)
			tp.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 17 Throttle position: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (tp *ThrottlePosition) GetLastValue() string {
	return tp.lastValue
}

// Unit Returns the unit of the pid, for the ThrottlePosition the unit is percentage
func (tp *ThrottlePosition) Unit() string {
	return "%"
}

// GetPid Returns the pid hex as integer
func (tp *ThrottlePosition) GetPid() int {
	return 0x11
}

// GetService Returns the service hex as integer
func (tp *ThrottlePosition) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (tp *ThrottlePosition) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", tp.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (tp *ThrottlePosition) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", tp.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (tp *ThrottlePosition) GetIsSupported() bool {
	return tp.IsSupported
}
