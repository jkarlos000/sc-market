package client

import (
	"context"
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/user/internal/core/domain"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/delivery/grpc/proto"
)

type UserClient	struct {
	client proto.UserServiceClient
	logger	hclog.Logger
}

func NewUserClient(c proto.UserServiceClient) *UserClient {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "UserClient",
		Level: hclog.LevelFromString("DEBUG"),
	})
	return &UserClient{
		client: c,
		logger: appLogger,
	}
}

func (uc *UserClient) parseToGRPC(user *domain.User) *proto.User {
	locked := *user.Locked
	userGrpc := &proto.User{
		ID:        uint32(user.ID),
		Email:     user.Email,
		Names:     user.Names,
		LastNames: user.LastNames,
		BirthDate: user.BirthDate,
		Gender:    user.Gender,
		Address:   user.Address,
		Reference: user.Reference,
		CI:        user.CI,
		Telephone: user.Telephone,
		LastIP:    user.LastIP,
		Locked:    locked,
	}
	return userGrpc
}

func (uc *UserClient) parseToUser(userGrpc *proto.User) *domain.User {
	locked := userGrpc.GetLocked()
	return &domain.User{
		ID:             uint(userGrpc.GetID()),
		Email:          userGrpc.GetEmail(),
		Names:          userGrpc.GetNames(),
		LastNames:      userGrpc.GetLastNames(),
		BirthDate:      userGrpc.GetBirthDate(),
		Gender:         userGrpc.GetGender(),
		Address:        userGrpc.GetAddress(),
		Reference:      userGrpc.GetReference(),
		CI:             userGrpc.GetCI(),
		Telephone:      userGrpc.GetTelephone(),
		JoinDate:       sql.NullTime{},
		LastIP:         userGrpc.GetLastIP(),
		Locked:         &locked,
	}
}

func (uc *UserClient) GetUser(id int) (*domain.User, error) {
	req := &proto.IdRequest{Id: uint32(id)}
	res, err := uc.client.Get(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al obtener el usuario", "err", err)
		return nil, err
	}
	return uc.parseToUser(res.GetUser()), nil
}
