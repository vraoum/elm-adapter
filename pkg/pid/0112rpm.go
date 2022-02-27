package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var rpmInstance = &rpm{}

// GetRpmInstance Return an instance of rpm
func GetRpmInstance() *rpm {
	return rpmInstance
}

type rpm struct {
	lastValue string // lastValue Last value received
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the rpm value, we must apply the following formula: (256 * a + b)/4
// Where a and b are the first and second parameters given
func (rpm *rpm) Convert(args []string) (string, error) {
	if len(args) == 2 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)
		b, decErr2 := strconv.ParseInt(args[1], 16, 64)
		if decErr1 == nil && decErr2 == nil {
			res := strconv.FormatInt((256*a+b)/4, 10)
			rpm.lastValue = res
			return res, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (rpm *rpm) GetLastValue() string {
	return rpm.lastValue
}

// Unit Returns the unit of the pid, for the rpm the unit is RPM as Rotation Per Minute
func (rpm *rpm) Unit() string {
	return "RPM"
}

// GetPid Returns the pid hex as integer
func (rpm *rpm) GetPid() int {
	return 0x0C
}

// GetService Returns the service hex as integer
func (rpm *rpm) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (rpm *rpm) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", rpm.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (rpm *rpm) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", rpm.GetService()))
}
