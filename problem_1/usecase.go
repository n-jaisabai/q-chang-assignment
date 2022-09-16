package problem_1

// Usecase represent the problem_1 usecases
type Usecase interface {
	SetEquation(equation func(x int) float64)
	FindNumberOfDatasetBySearch(dataset []interface{}, search []string) map[string]int
}
