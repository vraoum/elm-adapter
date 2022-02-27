package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Speed struct {
	lastValue string // lastValue Last value received
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the Speed value, we must parse the first argument as int
// Because a hex value on 2 characters can only be a value between 0 and 255,
// the max Speed that this method can return is 255
func (speed *Speed) Convert(args []string) (string, error) {
	a, decErr1 := strconv.ParseInt(args[0], 16, 64)
	if decErr1 == nil {
		res := strconv.FormatInt(a, 10)
		speed.lastValue = res
		return res, nil
	}
	return "", errors.New(fmt.Sprintf("Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (speed *Speed) GetLastValue() string {
	return speed.lastValue
}

// Unit Returns the unit of the pid, for the Speed the unit is KMH as kilometers per hour
func (speed *Speed) Unit() string {
	return "KMH"
}

// GetPid Returns the pid hex as integer
func (speed *Speed) GetPid() int {
	return 0x0D
}

// GetService Returns the service hex as integer
func (speed *Speed) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (speed *Speed) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", speed.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (speed *Speed) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", speed.GetService()))
}
