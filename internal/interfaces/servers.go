package interfaces

type ServerInterface interface {
	SyncStart()
	Stop()
}

var _ ServerInterface = &Servers{}

type Servers struct {
	Servers []ServerInterface
}

func (s *Servers) SyncStart() {
	for _, server := range s.Servers {
		server.SyncStart()
	}
}

func (s *Servers) Stop() {
	for _, server := range s.Servers {
		server.Stop()
	}
}
