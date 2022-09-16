package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/n-jaisabai/q-chang-assignment/models"
	"github.com/n-jaisabai/q-chang-assignment/problem_2"
)

// ChangeMoneyRequest represent the change money request struct
type ChangeMoneyRequest struct {
	Amount       float64 `json:"amount"`
	ProductPrice float64 `json:"product_price"`
}

// ChangeMoneyResponse represent the change money response struct
type ChangeMoneyResponse struct {
	Error  bool            `json:"error"`
	Change []*models.Money `json:"change"`
}

// ErrorResponse represent the response error struct
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Problem2Handler  represent the httphandler for problem_2
type Problem2Handler struct {
	Problem2Usecase problem_2.Usecase
}

// NewProblem2Handler will initialize the problem_2/ resources endpoint
func NewProblem2Handler(e *echo.Echo, u problem_2.Usecase) {
	handler := &Problem2Handler{
		Problem2Usecase: u,
	}
	e.POST("/problem-2", handler.ChangeMoney)
}

// ChangeMoney will return amount of coin
func (h *Problem2Handler) ChangeMoney(c echo.Context) error {
	var req ChangeMoneyRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ErrorResponse{Error: true, Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.Problem2Usecase.MakeChange(ctx, req.Amount, req.ProductPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: true, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ChangeMoneyResponse{
		Error:  false,
		Change: result,
	})
}
