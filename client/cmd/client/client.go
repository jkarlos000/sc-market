package client

import (
	"context"
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/client/internal/core/domain"
	"github.com/jkarlos000/sc-market/client/internal/infrastructure/delivery/grpc/proto"
	"google.golang.org/grpc"
)

type ClientCli struct {
client proto.ClientServiceClient
logger	hclog.Logger
}

func NewClientCli(c *grpc.ClientConn) *ClientCli {
	cc := proto.NewClientServiceClient(c)
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "ClientCli",
		Level: hclog.LevelFromString("DEBUG"),
	})
	return &ClientCli{
		client: cc,
		logger: appLogger,
	}
}

func (cc *ClientCli) parseToGRPC(client *domain.Client) *proto.Client {
	locked := *client.Locked
	clientGrpc := &proto.Client{
		ID:        uint32(client.ID),
		Email:     client.Email,
		Names:     client.Names,
		LastNames: client.LastNames,
		BirthDate: client.BirthDate,
		Gender:    client.Gender,
		Address:   client.Address,
		Reference: client.Reference,
		CI:        client.CI,
		Telephone: client.Telephone,
		LastIP:    client.LastIP,
		Locked:    locked,
	}
	return clientGrpc
}

func (cc *ClientCli) parseToClient(clientGrpc *proto.Client) *domain.Client {
	locked := clientGrpc.GetLocked()
	return &domain.Client{
		ID:             uint(clientGrpc.GetID()),
		Email:          clientGrpc.GetEmail(),
		Names:          clientGrpc.GetNames(),
		LastNames:      clientGrpc.GetLastNames(),
		BirthDate:      clientGrpc.GetBirthDate(),
		Gender:         clientGrpc.GetGender(),
		Address:        clientGrpc.GetAddress(),
		Reference:      clientGrpc.GetReference(),
		CI:             clientGrpc.GetCI(),
		Telephone:      clientGrpc.GetTelephone(),
		JoinDate:       sql.NullTime{},
		LastIP:         clientGrpc.GetLastIP(),
		Locked:         &locked,
	}
}

func (cc *ClientCli) listToGrpc(clients []*domain.Client) []*proto.Client {
	var listClientGrpc []*proto.Client
	for _, client := range clients {
		listClientGrpc = append(listClientGrpc, cc.parseToGRPC(client))
	}
	return listClientGrpc
}

func (cc *ClientCli) listToClient(clientsGrpc []*proto.Client) []*domain.Client {
	var listClient []*domain.Client
	for _, clientGrpc := range clientsGrpc {
		listClient = append(listClient, cc.parseToClient(clientGrpc))
	}
	return listClient
}

func (cc *ClientCli) GetClient(id int) (*domain.Client, error) {
	req := &proto.IdRequest{Id: uint32(id)}
	res, err := cc.client.Get(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al obtener el cliente", "err", err)
		return nil, err
	}
	return cc.parseToClient(res.GetClient()), nil
}

func (cc *ClientCli) CreateClient(client *domain.Client, password string) error {
	req := &proto.CreateRequest{
		Client:     cc.parseToGRPC(client),
		Password: password,
	}
	_, err := cc.client.Create(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al obtener el cliente", "err", err)
		return err
	}
	return nil
}

func (cc *ClientCli) ListClients() ([]*domain.Client, error) {
	req := &proto.ClientNilRequest{}
	resp, err := cc.client.List(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al obtener la Lista de clientes", "err", err)
		return nil, err
	}
	return cc.listToClient(resp.GetClients()), nil
}

func (cc *ClientCli) DeleteClient(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := cc.client.Delete(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al Eliminar el cliente", "err", err)
		return err
	}
	return nil
}

func (cc *ClientCli) DeleteClients(ids []int) error {
	var listId []uint32
	for _, value := range ids {
		listId = append(listId, uint32(value))
	}
	req := &proto.IdsRequest{Ids: listId}
	_, err := cc.client.Deletes(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al Eliminar los clientes", "err", err)
		return err
	}
	return nil
}

func (cc *ClientCli) UpdateClient(id int, client *domain.Client) error {
	req := &proto.UpdateRequest{
		Id:   uint32(id),
		Client: cc.parseToGRPC(client),
	}
	_, err := cc.client.Update(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al actualizar el cliente", "err", err)
		return err
	}
	return nil
}

func (cc *ClientCli) Login(email, password string) (bool, *domain.Client, error) {
	req := &proto.LoginRequest{
		Email:    email,
		Password: password,
	}
	resp, err := cc.client.Login(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error con el Login", "err", err)
		return false, nil, err
	}
	return resp.GetOk(), cc.parseToClient(resp.GetClient()), nil
}

func (cc *ClientCli) BanClient(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := cc.client.Ban(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al restringir al cliente", "err", err)
		return err
	}
	return nil
}

func (cc *ClientCli) UnbanClient(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := cc.client.Unban(context.Background(), req)
	if err != nil {
		cc.logger.Error("Ha ocurrido un error al quitar la restriccion al cliente", "err", err)
		return err
	}
	return nil
}
