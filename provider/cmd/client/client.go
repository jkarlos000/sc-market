package client

import (
	"context"
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/provider/internal/core/domain"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/delivery/grpc/proto"
)

type ProviderClient	struct {
client proto.ProviderServiceClient
logger	hclog.Logger
}

func NewProviderClient(c proto.ProviderServiceClient) *ProviderClient {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "ProviderClient",
		Level: hclog.LevelFromString("DEBUG"),
	})
	return &ProviderClient{
		client: c,
		logger: appLogger,
	}
}

func (pc *ProviderClient) parseToGRPC(provider *domain.Provider) *proto.Provider {
	locked := *provider.Locked
	providerGrpc := &proto.Provider{
		ID:        uint32(provider.ID),
		Email:     provider.Email,
		Names:     provider.Names,
		LastNames: provider.LastNames,
		BirthDate: provider.BirthDate,
		Gender:    provider.Gender,
		Address:   provider.Address,
		Reference: provider.Reference,
		CI:        provider.CI,
		Telephone: provider.Telephone,
		LastIP:    provider.LastIP,
		Locked:    locked,
	}
	return providerGrpc
}

func (pc *ProviderClient) parseToProvider(providerGrpc *proto.Provider) *domain.Provider {
	locked := providerGrpc.GetLocked()
	return &domain.Provider{
		ID:             uint(providerGrpc.GetID()),
		Email:          providerGrpc.GetEmail(),
		Names:          providerGrpc.GetNames(),
		LastNames:      providerGrpc.GetLastNames(),
		BirthDate:      providerGrpc.GetBirthDate(),
		Gender:         providerGrpc.GetGender(),
		Address:        providerGrpc.GetAddress(),
		Reference:      providerGrpc.GetReference(),
		CI:             providerGrpc.GetCI(),
		Telephone:      providerGrpc.GetTelephone(),
		JoinDate:       sql.NullTime{},
		LastIP:         providerGrpc.GetLastIP(),
		Locked:         &locked,
	}
}

func (pc *ProviderClient) listToGrpc(providers []*domain.Provider) []*proto.Provider {
	var listProviderGrpc []*proto.Provider
	for _, provider := range providers {
		listProviderGrpc = append(listProviderGrpc, pc.parseToGRPC(provider))
	}
	return listProviderGrpc
}

func (pc *ProviderClient) listToProvider(providersGrpc []*proto.Provider) []*domain.Provider {
	var listProvider []*domain.Provider
	for _, providerGrpc := range providersGrpc {
		listProvider = append(listProvider, pc.parseToProvider(providerGrpc))
	}
	return listProvider
}

func (pc *ProviderClient) GetProvider(id int) (*domain.Provider, error) {
	req := &proto.IdRequest{Id: uint32(id)}
	res, err := pc.client.Get(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al obtener el proveedor", "err", err)
		return nil, err
	}
	return pc.parseToProvider(res.GetProvider()), nil
}

func (pc *ProviderClient) CreateProvider(provider *domain.Provider, password string) error {
	req := &proto.CreateRequest{
		Provider:     pc.parseToGRPC(provider),
		Password: password,
	}
	_, err := pc.client.Create(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al obtener el proveedor", "err", err)
		return err
	}
	return nil
}

func (pc *ProviderClient) ListProviders() ([]*domain.Provider, error) {
	req := &proto.ProviderNilRequest{}
	resp, err := pc.client.List(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al obtener la Lista de proveedores", "err", err)
		return nil, err
	}
	return pc.listToProvider(resp.GetProviders()), nil
}

func (pc *ProviderClient) DeleteProvider(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := pc.client.Delete(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al Eliminar el proveedor", "err", err)
		return err
	}
	return nil
}

func (pc *ProviderClient) DeleteProviders(ids []int) error {
	var listId []uint32
	for _, value := range ids {
		listId = append(listId, uint32(value))
	}
	req := &proto.IdsRequest{Ids: listId}
	_, err := pc.client.Deletes(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al Eliminar los proveedores", "err", err)
		return err
	}
	return nil
}

func (pc *ProviderClient) UpdateProvider(id int, provider *domain.Provider) error {
	req := &proto.UpdateRequest{
		Id:   uint32(id),
		Provider: pc.parseToGRPC(provider),
	}
	_, err := pc.client.Update(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al actualizar el proveedor", "err", err)
		return err
	}
	return nil
}

func (pc *ProviderClient) Login(email, password string) (bool, *domain.Provider, error) {
	req := &proto.LoginRequest{
		Email:    email,
		Password: password,
	}
	resp, err := pc.client.Login(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error con el Login", "err", err)
		return false, nil, err
	}
	return resp.GetOk(), pc.parseToProvider(resp.GetProvider()), nil
}

func (pc *ProviderClient) BanProvider(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := pc.client.Ban(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al restringir al proveedor", "err", err)
		return err
	}
	return nil
}

func (pc *ProviderClient) UnbanProvider(id int) error {
	req := &proto.IdRequest{Id: uint32(id)}
	_, err := pc.client.Unban(context.Background(), req)
	if err != nil {
		pc.logger.Error("Ha ocurrido un error al quitar la restriccion al proveedor", "err", err)
		return err
	}
	return nil
}
