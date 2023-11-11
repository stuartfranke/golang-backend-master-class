package gapi

import (
	db "github.com/stuartfranke/golang-backend-master-class/db/sqlc"
	"github.com/stuartfranke/golang-backend-master-class/pb"
	"github.com/stuartfranke/golang-backend-master-class/token"
	"github.com/stuartfranke/golang-backend-master-class/util"
	"github.com/stuartfranke/golang-backend-master-class/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker := token.NewPasetoMaker(config.TokenSymmetricKey)
	//if err != nil {
	//	return nil, fmt.Errorf("cannot create token maker: %w", err)
	//}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
