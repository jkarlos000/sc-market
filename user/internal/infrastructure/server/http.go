package server

import (
	"context"
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/user/internal/core/domain"

	"github.com/jkarlos000/sc-market/user/internal/core/ports"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/delivery/grpc/proto"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	svc ports.UsersService
	l	hclog.Logger
}

func NewServer(svc ports.UsersService, l hclog.Logger)	*Server  {
	return &Server{
		svc: svc,
		l: l,
	}
}

func (s *Server) parseToGRPC(user *domain.User) *proto.User {
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

func (s *Server) parseToUser(userGrpc *proto.User) *domain.User {
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

func (s *Server) listToGrpc(users []*domain.User) []*proto.User {
	var listUserGrpc []*proto.User
	for _, user := range users {
		listUserGrpc = append(listUserGrpc, s.parseToGRPC(user))
	}
	return listUserGrpc
}

func (s *Server) listToUser(usersGrpc []*proto.User) []*domain.User {
	var listUser []*domain.User
	for _, userGrpc := range usersGrpc {
		listUser = append(listUser, s.parseToUser(userGrpc))
	}
	return listUser
}

func (s *Server) Get(ctx context.Context, request *proto.IdRequest) (*proto.UserResponse, error) {
	id := request.GetId()
	user, err := s.svc.Get(int(id))
	if err != nil {
		s.l.Error("Get: Ha ocurrido un error al obtener el usuario por ID","input", id, "error", err)
		return &proto.UserResponse{User: nil}, err
	}
	s.l.Info("Get: Se ha obtenido exitosamente el usuario con ID:", "input", id)
	return &proto.UserResponse{User: s.parseToGRPC(user)}, nil
}

func (s *Server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.MessageResponse, error) {
	user := s.parseToUser(request.GetUser())
	password := request.GetPassword()
	err := s.svc.Create(user, password)
	if err != nil {
		s.l.Error("Create: Ha ocurrido un error al crear el usuario", "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Create: se ha registrado satisfactoriamente el usuario", "input", user.Email)
	return &proto.MessageResponse{Message: "Usuario creado exitosamente"}, nil
}

func (s *Server) List(ctx context.Context, request *proto.UserNilRequest) (*proto.UsersResponse, error) {
	users, err := s.svc.List()
	if err != nil {
		s.l.Error("List: Ha ocurrido un error al Listar los usuarios", "error", err)
		return &proto.UsersResponse{Users: nil}, err
	}
	s.l.Info("List: se ha Listado un total de", "cantidad", len(users))
	return &proto.UsersResponse{Users: s.listToGrpc(users)}, nil
}

func (s *Server) Delete(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := request.GetId()
	err := s.svc.Delete(int(id))
	if err != nil {
		s.l.Error("Delete: Ha ocurrido un error al eliminar el usuario por ID","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Delete: Se ha eliminado exitosamente el usuario con ID:", "input", id)
	return &proto.MessageResponse{Message: "Usuario eliminado correctamente"}, nil
}

func (s *Server) Deletes(ctx context.Context, request *proto.IdsRequest) (*proto.MessageResponse, error) {
	ids := request.GetIds()
	for _, id := range ids {
		err := s.svc.Delete(int(id))
		if err != nil {
			s.l.Error("Deletes: Ha ocurrido un error al eliminar los usuarios","input", id, "error", err)
			return &proto.MessageResponse{Message: ""}, err
		}
	}
	s.l.Info("Deletes: Se ha eliminado exitosamente todos los usuarios")
	return &proto.MessageResponse{Message: "Usuarios eliminados"}, nil
}

func (s *Server) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.MessageResponse, error) {
	id := request.GetId()
	user := s.parseToUser(request.GetUser())
	err := s.svc.Update(int(id), user)
	if err != nil {
		s.l.Error("Update (Deprecated): Ha ocurrido un error al actualizar el usuario con ID","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Update (Deprecated): Se ha actualizado exitosamente el usuario con ID", "input", id)
	return &proto.MessageResponse{Message: "Usuario actualizado"}, nil
}

func (s *Server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	email := request.GetEmail()
	password := request.GetPassword()
	ok, user, err := s.svc.Login(email, password)
	if err != nil {
		s.l.Error("Login: Ha ocurrido un error al realizar Login","email", email, "error", err)
		return &proto.LoginResponse{
			Ok:   ok,
			User: nil,
		}, err
	}
	s.l.Info("Login: Se ha iniciado sesión correctamente con", "email", email)
	return &proto.LoginResponse{
		Ok:   ok,
		User: s.parseToGRPC(user),
	}, nil
}

func (s *Server) Ban(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := int(request.GetId())
	err := s.svc.Ban(id)
	if err != nil {
		s.l.Error("Ban: No se ha podido restringir al usuario","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Ban: Se ha restringido exitosamente al usuario", "input", id)
	return &proto.MessageResponse{Message: "Usuario baneado"}, nil
}

func (s *Server) Unban(ctx context.Context, request *proto.IdRequest) (*proto.MessageResponse, error) {
	id := int(request.GetId())
	err := s.svc.Unban(id)
	if err != nil {
		s.l.Error("Ban: No se ha podido quitar la restriccion al usuario","input", id, "error", err)
		return &proto.MessageResponse{Message: ""}, err
	}
	s.l.Info("Ban: Se ha quitado la restricción exitosamente al usuario", "input", id)
	return &proto.MessageResponse{Message: "Usuario desbaneado"}, nil
}

