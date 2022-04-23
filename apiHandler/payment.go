package apihandler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type createPaymentRequest struct {
	Amount string `json:"amount" binding:"required,min=1"`
	LoanID int64  `json:"loan_id" binding:"min=1" `
}
type paymentResponse struct {
	ID            int64     `json:"id"`
	LoanID        int64     `json:"loan_id"`
	UserID        int64     `json:"user_id"`
	Amount        string    `json:"amount"`
	CreatedBy     string    `json:"created_by"`
	LastUpdatedBy string    `json:"last_updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
}

func newPaymentResponse(loan db.Payment) paymentResponse {

	return paymentResponse{
		ID:            loan.ID,
		LoanID:        loan.LoanID,
		UserID:        loan.UserID,
		Amount:        loan.Amount,
		CreatedBy:     loan.CreatedBy,
		CreatedAt:     loan.CreatedAt,
		LastUpdatedBy: loan.LastUpdatedBy,
		LastUpdatedAt: loan.UpdatedAt,
	}

}

//swagger:route POST  /loans  createLoan
func (server *Server) createPayment(ctx *gin.Context) {
	var req createPaymentRequest

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreatePaymentParams{
		Amount:        req.Amount,
		LoanID:        req.LoanID,
		CreatedBy:     authPayload.Username,
		LastUpdatedBy: authPayload.Username,
		IpFrom:        ctx.Request.RemoteAddr,
		UserAgent:     ctx.Request.UserAgent(),
	}

	loan, err := server.store.CreatePaymentTerms(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newPaymentResponse(loan))
}

type getPaymentRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPayment(ctx *gin.Context) {
	var req getPaymentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	loan, err := server.store.GetPayment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newPaymentResponse(loan))
}

type listPaymentRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPayment(ctx *gin.Context) {
	var req listPaymentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListPaymentParams{
		// Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	payments, err := server.store.ListPayment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payments)
}
