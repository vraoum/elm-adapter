package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// EngineLoad The engine load describe the amount of power being made by the engine compare to the maximum amount of power
// it can make at the same rpm
type EngineLoad struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the EngineLoad value, we must apply the following formula: (100 * a) / 255
// Where a is the first parameters given
func (el *EngineLoad) Convert(args []string) (string, error) {
	if len(args) == 1 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)

		if decErr1 == nil {
			res := strconv.FormatInt((100*a)/255, 10)
			el.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 04 Engine load: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (el *EngineLoad) GetLastValue() string {
	return el.lastValue
}

// Unit Returns the unit of the pid, for the EngineLoad the unit is percentage
func (el *EngineLoad) Unit() string {
	return "%"
}

// GetPid Returns the pid hex as integer
func (el *EngineLoad) GetPid() int {
	return 0x04
}

// GetService Returns the service hex as integer
func (el *EngineLoad) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (el *EngineLoad) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", el.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (el *EngineLoad) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", el.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (el *EngineLoad) GetIsSupported() bool {
	return el.IsSupported
}
