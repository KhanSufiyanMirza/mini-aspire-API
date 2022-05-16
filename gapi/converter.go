package gapi

import (
	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/ma"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *ma.User {
	return &ma.User{
		Id:                user.ID,
		Name:              user.Name,
		Mobile:            user.Mobile.String,
		Address:           user.Address.String,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedBy:         user.CreatedBy,
		LastUpdatedBy:     user.LastUpdatedBy,
		CreatedAt:         timestamppb.New(user.CreatedAt),
		LastUpdatedAt:     timestamppb.New(user.UpdatedAt),
	}
}
