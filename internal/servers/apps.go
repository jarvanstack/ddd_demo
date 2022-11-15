package servers

import (
	"ddd_demo/internal/user"
)

type Apps struct {
	UserApp user.UserAppInterface
}

func NewApps(repos *Repos) *Apps {
	return &Apps{
		UserApp: user.NewUserApp(repos.UserRepo, repos.AuthRepo, repos.BillRepo),
	}
}
