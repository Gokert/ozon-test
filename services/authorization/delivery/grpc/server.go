package grpc

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"ozon-test/configs"
	"ozon-test/services/authorization/delivery/proto"
	"ozon-test/services/authorization/repository/profile"
	"ozon-test/services/authorization/repository/session"
)

type authGrpc struct {
	grpcServ *grpc.Server
	lg       *logrus.Logger
}

type server struct {
	proto.UnimplementedAuthorizationServer
	users    profile.IRepository
	sessions session.ISessionRepo
	lg       *logrus.Logger
}

func NewServer(l *logrus.Logger) (*authGrpc, error) {
	config, err := configs.GetAuthPsxConfig()
	if err != nil {
		return nil, fmt.Errorf("get auth psx config error: %s", err.Error())
	}

	cfgSession, err := configs.GetRedisConfig()
	if err != nil {
		return nil, fmt.Errorf("get config sessions error: %s", err.Error())
	}

	sessions, err := session.GetAuthRepo(cfgSession)
	if err != nil {
		return nil, fmt.Errorf("get sessions repo error: %s", err.Error())
	}

	users, err := profile.GetPsxRepo(config, l)
	if err != nil {
		return nil, fmt.Errorf("get psx error: %s", err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterAuthorizationServer(s, &server{
		lg:       l,
		users:    users,
		sessions: sessions,
	})

	return &authGrpc{grpcServ: s, lg: l}, nil
}

func (s *authGrpc) ListenAndServeGrpc(grpcCfg *configs.GrpcConfig) error {
	listen, err := net.Listen(grpcCfg.ConnectionType, ":"+grpcCfg.Port)
	if err != nil {
		return fmt.Errorf("listen auth grpc error: %s", err.Error())
	}

	err = s.grpcServ.Serve(listen)
	if err != nil {
		return fmt.Errorf("serve error: %s", err.Error())
	}

	return nil
}

func (s *server) GetId(ctx context.Context, req *proto.FindIdRequest) (*proto.FindIdResponse, error) {
	login, err := s.sessions.GetUserLogin(ctx, req.Sid)
	if err != nil {
		return nil, fmt.Errorf("get user login error: %s", err.Error())
	}

	id, err := s.users.GetUserId(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("get user id error: %s", err.Error())
	}

	return &proto.FindIdResponse{Value: id}, nil
}

func (s *server) GetAuthorizationStatus(ctx context.Context, req *proto.AuthorizationCheckRequest) (*proto.AuthorizationCheckResponse, error) {
	status, err := s.sessions.CheckActiveSession(ctx, req.Sid)
	if err != nil {
		return nil, fmt.Errorf("check auth status error: %s", err.Error())
	}
	return &proto.AuthorizationCheckResponse{
		Status: status,
	}, nil
}

func (s *server) GetRole(ctx context.Context, req *proto.RoleRequest) (*proto.RoleResponse, error) {
	role, err := s.users.GetRole(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("check role error: %s", err.Error())
	}

	return &proto.RoleResponse{Role: role}, nil
}

func (s *server) GetUserName(ctx context.Context, req *proto.UserItemRequest) (*proto.UserItemResponse, error) {
	name, err := s.users.GetUserName(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("get login error: %s", err.Error())
	}

	return &proto.UserItemResponse{Name: name}, nil
}
