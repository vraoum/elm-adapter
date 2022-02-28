package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Supported0120 The engine load describe the amount of power being made by the engine compare to the maximum amount of power
// it can make at the same rpm
type Supported0120 struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the Supported0120 value, we must apply the following formula: (100 * a) / 255
// Where a is the first parameters given
func (s *Supported0120) Convert(args []string) (string, error) {
	a, err := strconv.ParseInt(strings.Join(args, ""), 16, 64)
	if err == nil {
		res := strconv.FormatInt(a, 2)
		s.lastValue = res
		return res, nil
	}

	return "", errors.New(fmt.Sprintf("01 00 Supported 01 - 20: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (s *Supported0120) GetLastValue() string {
	return s.lastValue
}

// Unit Returns the unit of the pid, for the Supported0120 the unit is percentage
func (s *Supported0120) Unit() string {
	return ""
}

// GetPid Returns the pid hex as integer
func (s *Supported0120) GetPid() int {
	return 0x00
}

// GetService Returns the service hex as integer
func (s *Supported0120) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (s *Supported0120) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", s.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (s *Supported0120) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", s.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (s *Supported0120) GetIsSupported() bool {
	return s.IsSupported
}
