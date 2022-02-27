package serial

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/vraoum/elm-adapter/pkg/pid"
	"time"
)

type serialConnection struct {
	p        *serial.Port
	services map[int]*pid.Service
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

func (sc *serialConnection) GetPid(mode int, pid int) (pid.Pid, error) {
	if s, found := sc.services[mode]; found {
		if p, found := s.Pids[pid]; found {
			return p, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Could not find pid for serivce %d and pid %d", mode, pid))
}

func (sc *serialConnection) FindPid(p pid.Pid) (pid.Pid, error) {
	return sc.GetPid(p.GetService(), p.GetPid())
}

func OpenSerial(com string) *serialConnection {
	logrus.Infof("Attempting to connect to serial %s ...", com)
	c := &serial.Config{Name: com, Baud: 115200, ReadTimeout: time.Millisecond * 10}
	s, err := serial.OpenPort(c)

	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("connected")

	sc := &serialConnection{
		p:        s,
		services: make(map[int]*pid.Service),
	}

	sc.services[0x01] = pid.NewService(0x01)
	sc.services[0x09] = pid.NewService(0x09)

	return sc
}
