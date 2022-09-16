package usecase_test

import (
	"math"
	"testing"

	ucase "github.com/n-jaisabai/q-chang-assignment/problem_1/usecase"
	"github.com/stretchr/testify/assert"
)

func TestFindNumberOfDatasetBySearch(t *testing.T) {
	mockDataset := []interface{}{1, "X", 8, 17, "Y", "Z", 78, "W"}
	mockSearch := []string{"X", "Y", "Z", "W"}
	mockEquation := func(x int) float64 {
		return (1.0*math.Pow(float64(x), 3)+3.0*math.Pow(float64(x), 2)-4.0*float64(x))/6.0 + 1
	}

	u := ucase.NewProblem1Usecase()

	u.SetEquation(mockEquation)
	actual := u.FindNumberOfDatasetBySearch(mockDataset, mockSearch)

	assert := assert.New(t)

	// assert equality
	expected := map[string]int{
		"X": 3,
		"Y": 31,
		"Z": 51,
		"W": 113,
	}

	assert.Equal(expected, actual, "they should be equal")
}
