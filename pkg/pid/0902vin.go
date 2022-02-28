package pid

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Vin struct {
	lastValue string // lastValue Last value received
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the Vin value, we must parse all arguments as ascii character
func (vin *Vin) Convert(args []string) (string, error) {
	res, err := hex.DecodeString(strings.Join(args, ""))
	vin.lastValue = string(res)

	return string(res), err
}

// GetLastValue Returns the last value received
func (vin *Vin) GetLastValue() string {
	return vin.lastValue
}

// Unit Returns the unit of the pid, the Vin has no unit since it is an identifier
func (vin *Vin) Unit() string {
	return ""
}

// GetPid Returns the pid hex as integer
func (vin *Vin) GetPid() int {
	return 0x02
}

// GetService Returns the service hex as integer
func (vin *Vin) GetService() int {
	return 0x09
}

// GetStringPid Format GetPid as string in the format: %02x
func (vin *Vin) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", vin.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (vin *Vin) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", vin.GetService()))
}
