package http

import (
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/n-jaisabai/q-chang-assignment/problem_1"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Problem1Handler  represent the httphandler for problem_1
type Problem1Handler struct {
	Problem1Usecase problem_1.Usecase
}

// NewProblem1Handler will initialize the problem_1/ resources endpoint
func NewProblem1Handler(e *echo.Echo, u problem_1.Usecase) {
	handler := &Problem1Handler{
		Problem1Usecase: u,
	}
	e.GET("/problem-1", handler.GetProblem1)
}

// GetProblem1 will return value of dataset by searhing
func (h *Problem1Handler) GetProblem1(c echo.Context) error {
	dataset := []interface{}{1, "X", 8, 17, "Y", "Z", 78, 113}
	search := []string{"X", "Y", "Z"}
	equation := func(x int) float64 {
		return (1.0*math.Pow(float64(x), 3)+3.0*math.Pow(float64(x), 2)-4.0*float64(x))/6.0 + 1
	}

	h.Problem1Usecase.SetEquation(equation)
	result := h.Problem1Usecase.FindNumberOfDatasetBySearch(dataset, search)

	return c.JSON(http.StatusOK, result)
}
