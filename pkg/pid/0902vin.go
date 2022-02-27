package pid

import (
	"encoding/hex"
	"fmt"
	"strings"
)

var vinInstance = &vin{}

// GetVinInstance Return an instance of vin
func GetVinInstance() *vin {
	return vinInstance
}

type vin struct {
	lastValue string // lastValue Last value received
}

// Convert Transform a list of arguments to a value and sets the last value.
// In order to get the vin value, we must parse all arguments as ascii character
func (vin *vin) Convert(args []string) (string, error) {
	res, err := hex.DecodeString(strings.Join(args, ""))
	vin.lastValue = string(res)
	return string(res), err
}

// GetLastValue Returns the last value received
func (vin *vin) GetLastValue() string {
	return vin.lastValue
}

// Unit Returns the unit of the pid, the vin has no unit since it is an identifier
func (vin *vin) Unit() string {
	return ""
}

// GetPid Returns the pid hex as integer
func (vin *vin) GetPid() int {
	return 0x02
}

// GetService Returns the service hex as integer
func (vin *vin) GetService() int {
	return 0x09
}

// GetStringPid Format GetPid as string in the format: %02x
func (vin *vin) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", vin.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (vin *vin) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", vin.GetService()))
}
