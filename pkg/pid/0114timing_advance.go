package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TimingAdvance struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the TimingAdvance value, we must apply the following formula: (a / 2) - 40
// Where a is the first parameters given
// the range of TimingAdvance that this method can return is -64 to 63.5
func (ta *TimingAdvance) Convert(args []string) (string, error) {
	a, decErr1 := strconv.ParseInt(args[0], 16, 64)

	if decErr1 == nil {
		res := strconv.FormatInt((a/2)-40, 10)
		ta.lastValue = res

		return res, nil
	}

	return "", errors.New(fmt.Sprintf("01 14 Timing advance: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (ta *TimingAdvance) GetLastValue() string {
	return ta.lastValue
}

// Unit Returns the unit of the pid, for the TimingAdvance the unit is ° as degree before top dead center
func (ta *TimingAdvance) Unit() string {
	return "°"
}

// GetPid Returns the pid hex as integer
func (ta *TimingAdvance) GetPid() int {
	return 0x0E
}

// GetService Returns the service hex as integer
func (ta *TimingAdvance) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (ta *TimingAdvance) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", ta.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (ta *TimingAdvance) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", ta.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (ta *TimingAdvance) GetIsSupported() bool {
	return ta.IsSupported
}
