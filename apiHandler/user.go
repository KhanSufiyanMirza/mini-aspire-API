package apihandler

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Name          string `json:"name" binding:"required"`
	Mobile        string `json:"mobile" binding:"required"`
	Email         string `json:"email" binding:"email"`
	Password      string `json:"hashed_password" binding:"required,min=6"`
	CreatedBy     string `json:"created_by" binding:"required"`
	LastUpdatedBy string `json:"last_updated_by" binding:"required"`
}
type userResponse struct {
	ID                int64          `json:"id"`
	Name              string         `json:"name"`
	Mobile            sql.NullString `json:"mobile"`
	Address           sql.NullString `json:"address"`
	Email             string         `json:"email"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedBy         string         `json:"created_by"`
	LastUpdatedBy     string         `json:"last_updated_by"`
	CreatedAt         time.Time      `json:"created_at"`
	LastUpdatedAt     time.Time      `json:"last_updated_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:                user.ID,
		Name:              user.Name,
		Mobile:            user.Mobile,
		Address:           user.Address,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedBy:         user.CreatedBy,
		LastUpdatedBy:     user.LastUpdatedBy,
		CreatedAt:         user.CreatedAt,
		LastUpdatedAt:     user.UpdatedAt,
	}
}

//swagger:route POST  /users users createUser
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Name:           req.Name,
		Mobile:         sql.NullString{String: req.Mobile, Valid: len(strings.TrimSpace(req.Mobile)) > 0},
		Email:          req.Email,
		CreatedBy:      req.CreatedBy,
		LastUpdatedBy:  req.LastUpdatedBy,
		IpFrom:         ctx.Request.RemoteAddr,
		UserAgent:      ctx.Request.UserAgent(),
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if account.Owner != authPayload.Username {
	// 	err := errors.New("account doesn't belong to the authenticated user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, account)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListUserParams{
		// Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
}
type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUserByEmail(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		user.Email,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}
