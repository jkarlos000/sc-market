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

func (uc *UserClient) listToGrpc(users []*domain.User) []*proto.User {
	var listUserGrpc []*proto.User
	for _, user := range users {
		listUserGrpc = append(listUserGrpc, uc.parseToGRPC(user))
	}
	return listUserGrpc
}

func (uc *UserClient) listToUser(usersGrpc []*proto.User) []*domain.User {
	var listUser []*domain.User
	for _, userGrpc := range usersGrpc {
		listUser = append(listUser, uc.parseToUser(userGrpc))
	}
	return listUser
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

func (uc *UserClient) CreateUser(user *domain.User, password string) error {
	req := &proto.CreateRequest{
		User:     uc.parseToGRPC(user),
		Password: password,
	}
	_, err := uc.client.Create(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al obtener el usuario", "err", err)
		return err
	}
	return nil
}

func (uc *UserClient) ListUsers() ([]*domain.User, error) {
	req := &proto.UserNilRequest{}
	resp, err := uc.client.List(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al obtener la Lista de usuarios", "err", err)
		return nil, err
	}
	return uc.listToUser(resp.GetUsers()), nil
}

func (uc *UserClient) DeleteUser(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := uc.client.Delete(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al Eliminar el usuario", "err", err)
		return err
	}
	return nil
}

func (uc *UserClient) DeleteUsers(ids []int) error {
	var listId []uint32
	for _, value := range ids {
		listId = append(listId, uint32(value))
	}
	req := &proto.IdsRequest{Ids: listId}
	_, err := uc.client.Deletes(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al Eliminar los usuarios", "err", err)
		return err
	}
	return nil
}

func (uc *UserClient) UpdateUser(id int, user *domain.User) error {
	req := &proto.UpdateRequest{
		Id:   uint32(id),
		User: uc.parseToGRPC(user),
	}
	_, err := uc.client.Update(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al actualizar el usuario", "err", err)
		return err
	}
	return nil
}

func (uc *UserClient) Login(email, password string) (bool, *domain.User, error) {
	req := &proto.LoginRequest{
		Email:    email,
		Password: password,
	}
	resp, err := uc.client.Login(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error con el Login", "err", err)
		return false, nil, err
	}
	return resp.GetOk(), uc.parseToUser(resp.GetUser()), nil
}

func (uc *UserClient) BanUser(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := uc.client.Ban(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al restringir al usuario", "err", err)
		return err
	}
	return nil
}

func (uc *UserClient) UnbanUser(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := uc.client.Unban(context.Background(), req)
	if err != nil {
		uc.logger.Error("Ha ocurrido un error al quitar la restriccion al usuario", "err", err)
		return err
	}
	return nil
}
