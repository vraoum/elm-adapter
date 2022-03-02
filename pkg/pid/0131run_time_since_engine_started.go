package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// RunTimeSinceEngineStarted Define the position of the throttle body inside the engine
type RunTimeSinceEngineStarted struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the RunTimeSinceEngineStarted value, we must apply the following formula: 256 * a + b
// Where a and b are the first and second parameters given
// the range of RunTimeSinceEngineStarted that this method can return is 0 to 100
func (rtses *RunTimeSinceEngineStarted) Convert(args []string) (string, error) {
	if len(args) == 2 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)
		b, decErr2 := strconv.ParseInt(args[1], 16, 64)

		if decErr1 == nil && decErr2 == nil {
			res := strconv.FormatInt(256*a+b, 10)
			rtses.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 31 Run time since engine started: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (rtses *RunTimeSinceEngineStarted) GetLastValue() string {
	return rtses.lastValue
}

// Unit Returns the unit of the pid, for the RunTimeSinceEngineStarted the unit is s for second
func (rtses *RunTimeSinceEngineStarted) Unit() string {
	return "s"
}

// GetPid Returns the pid hex as integer
func (rtses *RunTimeSinceEngineStarted) GetPid() int {
	return 0x1F
}

// GetService Returns the service hex as integer
func (rtses *RunTimeSinceEngineStarted) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (rtses *RunTimeSinceEngineStarted) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", rtses.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (rtses *RunTimeSinceEngineStarted) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", rtses.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (rtses *RunTimeSinceEngineStarted) GetIsSupported() bool {
	return rtses.IsSupported
}
