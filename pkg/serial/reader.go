package serial

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vraoum/elm-adapter/pkg/pid"
	"github.com/vraoum/elm-adapter/pkg/util"
	"strconv"
	"strings"
	"time"
)

func (sc *serialConnection) Read() {
	for {
		buf := make([]byte, 128)
		_, err := sc.p.Read(buf)
		start := time.Now()
		if err == nil {
			str := string(buf)
			str = strings.ReplaceAll(str, " ", "")
			lines := strings.Split(str, "\r")
			instruction := ""
			for _, line := range lines {
				if len(line) > 1 {
					// Check if char at position 1 is ":"
					if line[1] == 58 {
						instruction = instruction + line[2:]
					}
					// Check if first char is "4"
					if line[0] == 52 {
						instruction = line
					}
				}
			}

			if instruction != "" {
				foundPid, res, err := sc.parse(instruction)
				if err != nil {
					logrus.Error(err)
				} else {
					logrus.Debugf("%s resulted in %s%s took %s\n", instruction, res, foundPid.Unit(), time.Since(start))
				}
			}
		}
	}
}

// parse Read a hex string, get the pid concerned and set it's value
func (sc *serialConnection) parse(str string) (pid.Pid, string, error) {
	strSplit := util.SplitSize(str, 2)
	pidClass, err := sc.getPidFromString(strSplit)
	if err != nil {
		return nil, "", err
	}
	resp, err := pidClass.Convert(strSplit[2:])
	if err != nil {
		return pidClass, "", err
	}
	return pidClass, resp, nil
}

// getPidFromString returns the Pid according to serial response
// We can get the correct service thanks to the response sent by the OBD
// The first two bytes sent in response are the service + 0x40 and the Pid
// So depending on the service and the pid received we can return the correct Pid instance
func (sc *serialConnection) getPidFromString(strSplit []string) (pid.Pid, error) {
	if len(strSplit) >= 2 {
		mode, modeErr := strconv.ParseInt(strSplit[0], 16, 64)
		requestPid, pidErr := strconv.ParseInt(strSplit[1], 16, 64)

		if modeErr == nil && pidErr == nil {
			mode = mode - 0x40
			return sc.GetPid(int(mode), int(requestPid))
		}
	}

	return nil, errors.New(fmt.Sprintf("Malformated string %s", strings.Join(strSplit, " ")))
}
