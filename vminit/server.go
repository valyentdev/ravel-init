package vminit

import (
	"os/exec"

	"github.com/valyentdev/ravel-init/proto"
	"google.golang.org/grpc"
)

type server struct {
	cmd     *exec.Cmd
	updates chan struct{}
	status  *proto.InitStatus
	server  *grpc.Server

	config Config
}

var _ proto.InitServiceServer = (*server)(nil)

func newInitAPI(config Config, cmd *exec.Cmd) *server {
	grpcServer := grpc.NewServer()
	server := &server{
		updates: make(chan struct{}, 1),
		config:  config,
		cmd:     cmd,
		status:  &proto.InitStatus{},
		server:  grpcServer,
	}

	proto.RegisterInitServiceServer(grpcServer, server)

	return server
}
