package serial

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/vraoum/elm-adapter/pkg/pid"
	"time"
)

type Connection struct {
	p        *serial.Port
	services map[int]*Service
}

func (sc *Connection) Write(str string) error {
	_, err := sc.p.Write([]byte(str))

	if err != nil {
		logrus.Errorf("An error occurred while trying to write %s to serial: %s", str, err)
		return err
	}

	return nil
}

func (sc *Connection) AskPid(pid pid.Pid) error {
	if pid.GetIsSupported() {
		logrus.Debugf("Asking pid %s %s", pid.GetStringService(), pid.GetStringPid())
		return sc.Write(fmt.Sprintf("%s %s\r", pid.GetStringService(), pid.GetStringPid()))
	}
	logrus.Debugf("Asking pid %s %s that is not supported", pid.GetStringService(), pid.GetStringPid())
	return errors.New(fmt.Sprintf("PID %s %s not supported", pid.GetStringService(), pid.GetStringPid()))
}

func (sc *Connection) GetPid(mode int, pid int) (pid.Pid, error) {
	if s, found := sc.services[mode]; found {
		if p, found := s.Pids[pid]; found {
			return p, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Could not find pid for serivce %d and pid %d", mode, pid))
}

func (sc *Connection) FindPid(p pid.Pid) (pid.Pid, error) {
	return sc.GetPid(p.GetService(), p.GetPid())
}

func OpenSerial(com string) *Connection {
	logrus.Infof("Attempting to connect to serial %s ...", com)
	c := &serial.Config{Name: com, Baud: 115200, ReadTimeout: time.Millisecond * 10}
	s, err := serial.OpenPort(c)

	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("connected")

	sc := &Connection{
		p:        s,
		services: make(map[int]*Service),
	}

	go sc.Read()

	sc.services[0x01] = NewService(0x01)
	sc.services[0x09] = NewService(0x09)

	for _, service := range sc.services {
		service.Initialize(sc)
	}

	return sc
}
