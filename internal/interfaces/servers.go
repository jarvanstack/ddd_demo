package interfaces

type ServerInterface interface {
	AsyncStart()
	Stop()
}

var _ ServerInterface = &Servers{}

type Servers struct {
	Servers []ServerInterface
}

func (s *Servers) AsyncStart() {
	for _, server := range s.Servers {
		server.AsyncStart()
	}
}

func (s *Servers) Stop() {
	for _, server := range s.Servers {
		server.Stop()
	}
}
