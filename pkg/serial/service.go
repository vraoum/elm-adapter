package serial

import (
	"github.com/sirupsen/logrus"
	"github.com/vraoum/elm-adapter/pkg/pid"
	"time"
)

type Service struct {
	Pids          map[int]pid.Pid
	serviceNumber int
}

func NewService(serviceNumber int) *Service {
	s := &Service{
		Pids:          make(map[int]pid.Pid),
		serviceNumber: serviceNumber,
	}

	return s
}

func (s *Service) Initialize(sc *Connection) {
	switch s.serviceNumber {
	case 0x01:
		s0120 := &pid.Supported0132{IsSupported: true}
		s.Pids[0x00] = s0120
		_ = sc.AskPid(s0120)
		for s0120.GetLastValue() == "" {
			time.Sleep(100 * time.Millisecond)
		}
		logrus.Debug("Supported pid 01 to 20:", s0120.GetLastValue())
		s.Pids[0x04] = &pid.EngineLoad{IsSupported: s0120.GetLastValue()[3] == 49}
		s.Pids[0x05] = &pid.EngineCoolantTemperature{IsSupported: s0120.GetLastValue()[4] == 49}
		s.Pids[0x0A] = &pid.FuelPressure{IsSupported: s0120.GetLastValue()[9] == 49}
		s.Pids[0x0B] = &pid.IntakeManifoldAbsolutePressure{IsSupported: s0120.GetLastValue()[10] == 49}
		s.Pids[0x0C] = &pid.Rpm{IsSupported: s0120.GetLastValue()[11] == 49}
		s.Pids[0x0D] = &pid.Speed{IsSupported: s0120.GetLastValue()[12] == 49}
		s.Pids[0x0E] = &pid.TimingAdvance{IsSupported: s0120.GetLastValue()[13] == 49}
		s.Pids[0x0F] = &pid.IntakeAirTemperature{IsSupported: s0120.GetLastValue()[14] == 49}
		s.Pids[0x10] = &pid.MassAirFlowSensor{IsSupported: s0120.GetLastValue()[15] == 49}
		s.Pids[0x11] = &pid.ThrottlePosition{IsSupported: s0120.GetLastValue()[16] == 49}

	case 0x09:
		s.Pids[0x02] = &pid.Vin{}
	}
}
