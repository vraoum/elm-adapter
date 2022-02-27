package serial

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/vraoum/elm-adapter/pkg/pid"
	"time"
)

type serialConnection struct {
	p *serial.Port
}

func (sc *serialConnection) Write(str string) error {
	_, err := sc.p.Write([]byte(str))
	if err != nil {
		logrus.Errorf("An error occurred while trying to write %s to serial: %s", str, err)
		return err
	}
	return nil
}

func (sc *serialConnection) AskPid(pid pid.Pid) error {
	logrus.Debugf("Asking pid %s %s", pid.GetStringService(), pid.GetStringPid())
	return sc.Write(fmt.Sprintf("%s %s\r", pid.GetStringService(), pid.GetStringPid()))
}

func OpenSerial(com string) *serialConnection {
	logrus.Infof("Attempting to connect to serial %s ...", com)
	c := &serial.Config{Name: com, Baud: 115200, ReadTimeout: time.Millisecond * 10}
	s, err := serial.OpenPort(c)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("connected")
	return &serialConnection{
		p: s,
	}
}
