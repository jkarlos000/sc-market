package server

import (
	"context"
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/provider/internal/core/domain"
	"github.com/jkarlos000/sc-market/provider/internal/core/ports"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/delivery/grpc/proto"
)

type Server struct {
	proto.UnimplementedProviderServiceServer
	svc ports.ProvidersService
	l	hclog.Logger
}

func NewServer(svc ports.ProvidersService, l hclog.Logger)	*Server  {
	return &Server{
		svc: svc,
		l: l,
	}
}

func (s *Server) parseToGRPC(provider *domain.Provider) *proto.Provider {
	locked := *provider.Locked
	providerGrpc := &proto.Provider{
		ID:        uint32(provider.ID),
		Email:     provider.Email,
		Names:     provider.Names,
		LastNames: provider.LastNames,
		BirthDate: provider.BirthDate,
		BusinessName: provider.BusinessName,
		Nit: provider.Nit,
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

func (s *Server) parseToProvider(providerGrpc *proto.Provider) *domain.Provider {
	locked := providerGrpc.GetLocked()
	return &domain.Provider{
		ID:             uint(providerGrpc.GetID()),
		Email:          providerGrpc.GetEmail(),
		Names:          providerGrpc.GetNames(),
		LastNames:      providerGrpc.GetLastNames(),
		BirthDate:      providerGrpc.GetBirthDate(),
		BusinessName: providerGrpc.GetBusinessName(),
		Nit: providerGrpc.Nit,
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

func (s *Server) listToGrpc(providers []*domain.Provider) []*proto.Provider {
	var listProviderGrpc []*proto.Provider
	for _, provider := range providers {
		listProviderGrpc = append(listProviderGrpc, s.parseToGRPC(provider))
	}
	return listProviderGrpc
}

func (s *Server) listToProvider(providersGrpc []*proto.Provider) []*domain.Provider {
	var listProvider []*domain.Provider
	for _, providerGrpc := range providersGrpc {
		listProvider = append(listProvider, s.parseToProvider(providerGrpc))
	}
	return listProvider
}

func (s *Server) Get(ctx context.Context, request *proto.IdRequest) (*proto.ProviderResponse, error) {
	id := request.GetId()
	provider, err := s.svc.Get(int(id))
	if err != nil {
		s.l.Error("Get: Ha ocurrido un error al obtener el proveedor por ID","input", id, "error", err)
		return &proto.ProviderResponse{Provider: nil}, err
	}
	s.l.Info("Get: Se ha obtenido exitosamente el proveedor con ID:", "input", id)
	return &proto.ProviderResponse{Provider: s.parseToGRPC(provider)}, nil
}

func (s *Server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.MessageResponse, error) {
	provider := s.parseToProvider(request.GetProvider())
	password := request.GetPassword()
	err := s.svc.Create(provider, password)
	if err != nil {
		s.l.Error("Create: Ha ocurrido un error al crear el proveedor", "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Create: se ha registrado satisfactoriamente el proveedor", "input", provider.Email)
	return &proto.MessageResponse{Message: "Proveedor creado exitosamente"}, nil
}

func (s *Server) List(ctx context.Context, request *proto.ProviderNilRequest) (*proto.ProvidersResponse, error) {
	providers, err := s.svc.List()
	if err != nil {
		s.l.Error("List: Ha ocurrido un error al Listar los proveedors", "error", err)
		return &proto.ProvidersResponse{Providers: nil}, err
	}
	s.l.Info("List: se ha Listado un total de", "cantidad", len(providers))
	return &proto.ProvidersResponse{Providers: s.listToGrpc(providers)}, nil
}

func (s *Server) Delete(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := request.GetId()
	err := s.svc.Delete(int(id))
	if err != nil {
		s.l.Error("Delete: Ha ocurrido un error al eliminar el proveedor por ID","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Delete: Se ha eliminado exitosamente el proveedor con ID:", "input", id)
	return &proto.MessageResponse{Message: "Proveedor eliminado correctamente"}, nil
}

func (s *Server) Deletes(ctx context.Context, request *proto.IdsRequest) (*proto.MessageResponse, error) {
	ids := request.GetIds()
	for _, id := range ids {
		err := s.svc.Delete(int(id))
		if err != nil {
			s.l.Error("Deletes: Ha ocurrido un error al eliminar los proveedors","input", id, "error", err)
			return &proto.MessageResponse{Message: ""}, err
		}
	}
	s.l.Info("Deletes: Se ha eliminado exitosamente todos los proveedors")
	return &proto.MessageResponse{Message: "Proveedors eliminados"}, nil
}

func (s *Server) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.MessageResponse, error) {
	id := request.GetId()
	provider := s.parseToProvider(request.GetProvider())
	err := s.svc.Update(int(id), provider)
	if err != nil {
		s.l.Error("Update (Deprecated): Ha ocurrido un error al actualizar el proveedor con ID","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Update (Deprecated): Se ha actualizado exitosamente el proveedor con ID", "input", id)
	return &proto.MessageResponse{Message: "Proveedor actualizado"}, nil
}

func (s *Server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	email := request.GetEmail()
	password := request.GetPassword()
	ok, provider, err := s.svc.Login(email, password)
	if err != nil {
		s.l.Error("Login: Ha ocurrido un error al realizar Login","email", email, "error", err)
		return &proto.LoginResponse{
			Ok:   ok,
			Provider: nil,
		}, err
	}
	s.l.Info("Login: Se ha iniciado sesión correctamente con", "email", email)
	return &proto.LoginResponse{
		Ok:   ok,
		Provider: s.parseToGRPC(provider),
	}, nil
}

func (s *Server) Ban(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := int(request.GetId())
	err := s.svc.Ban(id)
	if err != nil {
		s.l.Error("Ban: No se ha podido restringir al proveedor","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Ban: Se ha restringido exitosamente al proveedor", "input", id)
	return &proto.MessageResponse{Message: "Proveedor baneado"}, nil
}

func (s *Server) Unban(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := int(request.GetId())
	err := s.svc.Unban(id)
	if err != nil {
		s.l.Error("Ban: No se ha podido quitar la restriccion al proveedor","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Ban: Se ha quitado la restricción exitosamente al proveedor", "input", id)
	return &proto.MessageResponse{Message: "Proveedor desbaneado"}, nil
}

