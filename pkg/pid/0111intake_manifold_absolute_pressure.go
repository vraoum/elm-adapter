package pid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// IntakeManifoldAbsolutePressure Define pressure in the intake manifold
type IntakeManifoldAbsolutePressure struct {
	lastValue   string // lastValue Last value received
	IsSupported bool
}

// Convert Transform a list of arguments to a value and sets the last value.
// the range of IntakeManifoldAbsolutePressure that this method can return is 0 to 255
func (imap *IntakeManifoldAbsolutePressure) Convert(args []string) (string, error) {
	if len(args) == 1 {
		a, decErr1 := strconv.ParseInt(args[0], 16, 64)

		if decErr1 == nil {
			res := strconv.FormatInt(a, 10)
			imap.lastValue = res

			return res, nil
		}
	}

	return "", errors.New(fmt.Sprintf("01 11 Intake manifold absolute pressure: Error while parsing with args %v", args))
}

// GetLastValue Returns the last value received
func (imap *IntakeManifoldAbsolutePressure) GetLastValue() string {
	return imap.lastValue
}

// Unit Returns the unit of the pid, for the FuelPressure the unit is kPa as kilopascal
func (imap *IntakeManifoldAbsolutePressure) Unit() string {
	return "kPa"
}

// GetPid Returns the pid hex as integer
func (imap *IntakeManifoldAbsolutePressure) GetPid() int {
	return 0x0B
}

// GetService Returns the service hex as integer
func (imap *IntakeManifoldAbsolutePressure) GetService() int {
	return 0x01
}

// GetStringPid Format GetPid as string in the format: %02x
func (imap *IntakeManifoldAbsolutePressure) GetStringPid() string {
	return strings.ToUpper(fmt.Sprintf("%02x", imap.GetPid()))
}

// GetStringService Format GetService as string in the format: %02x
func (imap *IntakeManifoldAbsolutePressure) GetStringService() string {
	return strings.ToUpper(fmt.Sprintf("%02x", imap.GetService()))
}

// GetIsSupported Returns whether the pid is supported or not
func (imap *IntakeManifoldAbsolutePressure) GetIsSupported() bool {
	return imap.IsSupported
}
