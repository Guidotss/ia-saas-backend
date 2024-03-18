package datasources

import (
	"github.com/Guidotss/ia-saas-backend.git/domain/dtos/auth"
)

type AuthDatasource interface {
	Login(userData auth.LoginDto) (string, error)
	Register(userData auth.RegisterDto) error
}
