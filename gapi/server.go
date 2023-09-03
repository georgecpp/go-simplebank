package gapi

import (
	"fmt"

	db "github.com/georgecpp/go-simplebank/db/sqlc"
	"github.com/georgecpp/go-simplebank/pb"
	"github.com/georgecpp/go-simplebank/token"
	"github.com/georgecpp/go-simplebank/util"
	"github.com/georgecpp/go-simplebank/worker"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store  db.Store
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server and setup routing
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor,
	}
	
	return server, nil
}