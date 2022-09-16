package usecase

import (
	"github.com/n-jaisabai/q-chang-assignment/problem_1"
	"github.com/thoas/go-funk"
)

type problem1Usecase struct {
	equation func(x int) float64
}

// NewProblem1Usecase will create new an problem1Usecase object representation of problem_1.Usecase interface
func NewProblem1Usecase() problem_1.Usecase {
	return &problem1Usecase{}
}

func (u *problem1Usecase) SetEquation(equation func(x int) float64) {
	u.equation = equation
}

// FindNumberOfDatasetBySearch implements problem_1.Usecase
func (u *problem1Usecase) FindNumberOfDatasetBySearch(dataset []interface{}, search []string) map[string]int {
	var result = make(map[string]int, len(search))
	for _, s := range search {
		index := funk.IndexOf(dataset, s)
		if index < 0 {
			result[s] = -1
		} else {
			result[s] = int(u.findNumberBySequence(index + 1))
		}
	}

	return result
}

func (u *problem1Usecase) findNumberBySequence(sequence int) float64 {
	return u.equation(sequence)
}
