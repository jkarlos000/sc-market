package server

import (
	"context"
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/client/internal/core/domain"
	"github.com/jkarlos000/sc-market/client/internal/core/ports"
	"github.com/jkarlos000/sc-market/client/internal/infrastructure/delivery/grpc/proto"
)

type Server struct {
	proto.UnimplementedClientServiceServer
	svc ports.ClientsService
	l	hclog.Logger
}

func NewServer(svc ports.ClientsService, l hclog.Logger)	*Server  {
	return &Server{
		svc: svc,
		l: l,
	}
}

func (s *Server) parseToGRPC(client *domain.Client) *proto.Client {
	locked := *client.Locked
	clientGrpc := &proto.Client{
		ID:        uint32(client.ID),
		Email:     client.Email,
		Names:     client.Names,
		LastNames: client.LastNames,
		BirthDate: client.BirthDate,
		BusinessName: client.BusinessName,
		Nit: client.Nit,
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

func (s *Server) parseToClient(clientGrpc *proto.Client) *domain.Client {
	locked := clientGrpc.GetLocked()
	return &domain.Client{
		ID:             uint(clientGrpc.GetID()),
		Email:          clientGrpc.GetEmail(),
		Names:          clientGrpc.GetNames(),
		LastNames:      clientGrpc.GetLastNames(),
		BirthDate:      clientGrpc.GetBirthDate(),
		BusinessName: clientGrpc.GetBusinessName(),
		Nit: clientGrpc.Nit,
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

func (s *Server) listToGrpc(clients []*domain.Client) []*proto.Client {
	var listClientGrpc []*proto.Client
	for _, client := range clients {
		listClientGrpc = append(listClientGrpc, s.parseToGRPC(client))
	}
	return listClientGrpc
}

func (s *Server) listToClient(clientsGrpc []*proto.Client) []*domain.Client {
	var listClient []*domain.Client
	for _, clientGrpc := range clientsGrpc {
		listClient = append(listClient, s.parseToClient(clientGrpc))
	}
	return listClient
}

func (s *Server) Get(ctx context.Context, request *proto.IdRequest) (*proto.ClientResponse, error) {
	id := request.GetId()
	client, err := s.svc.Get(int(id))
	if err != nil {
		s.l.Error("Get: Ha ocurrido un error al obtener el cliente por ID","input", id, "error", err)
		return &proto.ClientResponse{Client: nil}, err
	}
	s.l.Info("Get: Se ha obtenido exitosamente el cliente con ID:", "input", id)
	return &proto.ClientResponse{Client: s.parseToGRPC(client)}, nil
}

func (s *Server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.MessageResponse, error) {
	client := s.parseToClient(request.GetClient())
	password := request.GetPassword()
	err := s.svc.Create(client, password)
	if err != nil {
		s.l.Error("Create: Ha ocurrido un error al crear el cliente", "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Create: se ha registrado satisfactoriamente el cliente", "input", client.Email)
	return &proto.MessageResponse{Message: "Cliente creado exitosamente"}, nil
}

func (s *Server) List(ctx context.Context, request *proto.ClientNilRequest) (*proto.ClientsResponse, error) {
	clients, err := s.svc.List()
	if err != nil {
		s.l.Error("List: Ha ocurrido un error al Listar los clientes", "error", err)
		return &proto.ClientsResponse{Clients: nil}, err
	}
	s.l.Info("List: se ha Listado un total de", "cantidad", len(clients))
	return &proto.ClientsResponse{Clients: s.listToGrpc(clients)}, nil
}

func (s *Server) Delete(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := request.GetId()
	err := s.svc.Delete(int(id))
	if err != nil {
		s.l.Error("Delete: Ha ocurrido un error al eliminar el cliente por ID","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Delete: Se ha eliminado exitosamente el cliente con ID:", "input", id)
	return &proto.MessageResponse{Message: "Cliente eliminado correctamente"}, nil
}

func (s *Server) Deletes(ctx context.Context, request *proto.IdsRequest) (*proto.MessageResponse, error) {
	ids := request.GetIds()
	for _, id := range ids {
		err := s.svc.Delete(int(id))
		if err != nil {
			s.l.Error("Deletes: Ha ocurrido un error al eliminar los clientes","input", id, "error", err)
			return &proto.MessageResponse{Message: ""}, err
		}
	}
	s.l.Info("Deletes: Se ha eliminado exitosamente todos los clientes")
	return &proto.MessageResponse{Message: "Clientes eliminados"}, nil
}

func (s *Server) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.MessageResponse, error) {
	id := request.GetId()
	client := s.parseToClient(request.GetClient())
	err := s.svc.Update(int(id), client)
	if err != nil {
		s.l.Error("Update (Deprecated): Ha ocurrido un error al actualizar el cliente con ID","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Update (Deprecated): Se ha actualizado exitosamente el cliente con ID", "input", id)
	return &proto.MessageResponse{Message: "Cliente actualizado"}, nil
}

func (s *Server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	email := request.GetEmail()
	password := request.GetPassword()
	ok, client, err := s.svc.Login(email, password)
	if err != nil {
		s.l.Error("Login: Ha ocurrido un error al realizar Login","email", email, "error", err)
		return &proto.LoginResponse{
			Ok:   ok,
			Client: nil,
		}, err
	}
	s.l.Info("Login: Se ha iniciado sesión correctamente con", "email", email)
	return &proto.LoginResponse{
		Ok:   ok,
		Client: s.parseToGRPC(client),
	}, nil
}

func (s *Server) Ban(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := int(request.GetId())
	err := s.svc.Ban(id)
	if err != nil {
		s.l.Error("Ban: No se ha podido restringir al cliente","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Ban: Se ha restringido exitosamente al cliente", "input", id)
	return &proto.MessageResponse{Message: "Cliente baneado"}, nil
}

func (s *Server) Unban(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := int(request.GetId())
	err := s.svc.Unban(id)
	if err != nil {
		s.l.Error("Ban: No se ha podido quitar la restriccion al cliente","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Ban: Se ha quitado la restricción exitosamente al cliente", "input", id)
	return &proto.MessageResponse{Message: "Cliente desbaneado"}, nil
}

