package datasources

import (
	"github.com/Guidotss/ia-saas-backend.git/domain/dtos/auth"
)

type AuthDatasourceImpl struct {
}

func (a *AuthDatasourceImpl) Login(userData auth.LoginDto) (string, error) {
	return "", nil
}
