package gapi

import (
	db "github.com/stuartfranke/golang-backend-master-class/db/sqlc"
	"github.com/stuartfranke/golang-backend-master-class/pb"
	"github.com/stuartfranke/golang-backend-master-class/token"
	"github.com/stuartfranke/golang-backend-master-class/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker := token.NewPasetoMaker(config.TokenSymmetricKey)
	//if err != nil {
	//	return nil, fmt.Errorf("cannot create token maker: %v", err)
	//}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
