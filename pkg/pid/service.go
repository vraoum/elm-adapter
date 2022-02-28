package pid

type Service struct {
	Pids map[int]Pid
}

func NewService(serviceNumber int) *Service {
	s := &Service{
		Pids: make(map[int]Pid),
	}

	switch serviceNumber {
	case 0x01:
		s.Pids[0x04] = &EngineLoad{}
		s.Pids[0x0C] = &Rpm{}
		s.Pids[0x0D] = &Speed{}

	case 0x09:
		s.Pids[0x02] = &Vin{}
	}

	return s
}
