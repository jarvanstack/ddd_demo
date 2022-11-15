package servers

import (
	"ddd_demo/internal/user"
)

type Apps struct {
	UserApp user.UserAppInterface
}

func NewApps(apps *Repos) *Apps {
	return &Apps{
		UserApp: user.NewUserApp(apps.UserRepo, apps.AuthRepo),
	}
}
