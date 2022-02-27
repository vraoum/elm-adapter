package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var speedInstance = &speed{}

// GetSpeedInstance Return an instance of speed
func GetSpeedInstance() *speed {
	return speedInstance
}

type speed struct {
	lastValue string // lastValue Last value received
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the speed value, we must parse the first argument as int
// Because a hex value on 2 characters can only be a value between 0 and 255,
// the max speed that this method can return is 255
func (speed *speed) Convert(args []string) (string, error) {
	a, decErr1 := strconv.ParseInt(args[0], 16, 64)
	if decErr1 == nil {
		res := strconv.FormatInt(a, 10)
		speed.lastValue = res
		return res, nil
	}
	return "", errors.New(fmt.Sprintf("Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (speed *speed) GetLastValue() string {
	return speed.lastValue
}

// Unit Returns the unit of the pid, for the speed the unit is KMH as kilometers per hour
func (speed *speed) Unit() string {
	return "KMH"
}

// GetPid Returns the pid hex as integer
func (speed *speed) GetPid() int {
	return 0x0D
}

// GetService Returns the service hex as integer
func (speed *speed) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (speed *speed) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", speed.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (speed *speed) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", speed.GetService()))
}
