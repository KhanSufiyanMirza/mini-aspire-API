package gapi

import (
	"context"
	"database/sql"

	"github.com/KhanSufiyanMirza/mini-aspire-API/ma"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *ma.LoginUserRequest) (*ma.LoginUserResponse, error) {
	user, err := server.store.GetUserByEmail(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User not found  :%s", err)

		}
		return nil, status.Errorf(codes.Internal, "Failed to find user  :%s", err)

	}
	err = utils.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Please Enter Proper Details like Password  :%s", err)
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		user.Email,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create access token  :%s", err)

	}
	rsp := &ma.LoginUserResponse{
		AccessToken: accessToken,
		User:        convertUser(user),
	}

	return rsp, nil
}
