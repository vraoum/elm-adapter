package pid

type Pid interface {
	Convert(args []string) (string, error) // Convert Transform a list of arguments to a value and sets the last value
	GetLastValue() string                  // GetLastValue Returns the last value received
	Unit() string                          // Unit Returns the unit of the pid
	GetPid() int                           // GetPid Returns the pid hex as integer
	GetStringPid() string                  // GetStringPid Format GetPid as string in the format: %02x
	GetService() int                       // GetService Returns the service hex as integer
	GetStringService() string              // GetStringService Format GetService as string in the format: %02x
	GetIsSupported() bool                  // GetIsSupported Returns whether the pid is supported or not
}
