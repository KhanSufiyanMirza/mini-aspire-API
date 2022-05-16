package gapi

import (
	"fmt"

	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/ma"
	"github.com/KhanSufiyanMirza/mini-aspire-API/token"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
)

type Server struct {
	ma.UnimplementedMiniAspireServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

//NewServer Creates a new gRPC server.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("connot create token maker : %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokMaker,
	}

	return server, nil
}
