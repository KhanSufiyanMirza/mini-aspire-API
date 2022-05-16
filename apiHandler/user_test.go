package apihandler

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	mockdb "github.com/KhanSufiyanMirza/mini-aspire-API/db/mock"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32, false),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestGetUserAPI(t *testing.T) {
	user := createRandomUser()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	//build stubs
	store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).
		Times(1).Return(user, nil)

	//start test server and send request
	server := newTestServer(t, store)
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/users/%d", user.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	//check response
	require.Equal(t, http.StatusOK, recorder.Code)
}
func createRandomUser() db.User {
	hashedPassword, _ := utils.HashPassword(utils.RandomString(6, false))
	return db.User{
		ID:                utils.RandomInt(1, 1000),
		Name:              utils.RandomOwner(),
		Mobile:            sql.NullString{String: utils.RandomMobile(), Valid: true},
		Address:           sql.NullString{String: utils.RandomString(10, false), Valid: true},
		Email:             utils.RandomEmail(),
		HashedPassword:    hashedPassword,
		PasswordChangedAt: time.Now(),
		IsActive:          true,
		CreatedBy:         utils.RandomOwner(),
		CreatedAt:         time.Now(),
		LastUpdatedBy:     utils.RandomOwner(),
		UpdatedAt:         time.Now(),
		IpFrom:            "",
		UserAgent:         "",
	}
}
