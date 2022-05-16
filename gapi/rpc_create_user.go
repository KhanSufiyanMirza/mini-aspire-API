package gapi

import (
	"context"
	"database/sql"
	"strings"

	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/ma"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *ma.CreatUserRequest) (*ma.CreatUserResponse, error) {
	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password: %s", err)
	}
	p, _ := peer.FromContext(ctx)

	arg := db.CreateUserParams{
		Name:           req.GetName(),
		Mobile:         sql.NullString{String: req.GetMobile(), Valid: len(strings.TrimSpace(req.GetMobile())) > 0},
		Email:          req.GetEmail(),
		CreatedBy:      req.GetEmail(),
		LastUpdatedBy:  req.GetEmail(),
		IpFrom:         p.Addr.String(),
		UserAgent:      p.Addr.Network(), //Todo: in case of http how can i get this
		HashedPassword: hashedPassword,
		Address:        sql.NullString{String: req.GetAddress(), Valid: len(strings.TrimSpace(req.GetAddress())) > 0},
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "Email already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "Failed to create user: %s", err)
	}

	resp := &ma.CreatUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}
