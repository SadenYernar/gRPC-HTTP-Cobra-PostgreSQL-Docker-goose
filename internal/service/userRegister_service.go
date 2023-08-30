package service

import (
	"context"

	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/entity"
)

type IUserRegisterService interface {
	UserRegister(ctx context.Context, in entity.User) error
	//GetUserByPhone(ctx context.Context)
}
