package server

import (
	"context"
	"errors"
	"fmt"
	pb "gRPC-HTTP-Cobra-PostgreSQL-Docker-goose/api/proto/shop"
	"time"

	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/entity"
	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/logger"
	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/service"
)

type Server struct {
	pb.UnimplementedShopServer
	RegisterService   service.IUserRegisterService
	AddProductService service.IAddProduct
}

func (s *Server) UserRegisterServer(ctx context.Context, in *pb.UserRegisterRequest) (*pb.UserRegisterResponce, error) {
	logger.Info("Start working Register")

	UserInfo := &entity.User{
		Id:          in.GetId,
		FirstName:   in.GetFirstName,
		LastName:    in.GetLastName,
		Username:    in.GetUsername,
		Email:       in.GetEmail,
		Password:    in.GetPassword,
		Sex:         in.GetSex,
		City:        in.GetCity,
		DateBirth:   in.GetDateBirth,
		DateCreated: time.Now(),
	}

	if err := s.RegisterService.UserRegister(ctx, *UserInfo); err != nil {
		logger.Error("Register error: " + err.Error())

		if errors.Is(err, fmt.Errorf("AlreadyRegistered")) {
			return &pb.UserRegisterResponce{
				Error:   2,
				Message: "Уже зарегистрирован",
			}, nil
		}
		return nil, err
	}

	return &pb.UserRegisterResponce{
		Status: "OK",
	}, nil
}
