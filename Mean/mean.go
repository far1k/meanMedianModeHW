package Mean

import "errors"

func SumUp(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Please enter a non-empty slice")
	}
	sum := 0
	for i := range slice {
		sum += slice[i]
	}
	return sum, nil
}

func MeanFloat(slice []int) float64 {
	sumOfElements, _ := SumUp(slice)
	mean := float64(sumOfElements) / float64(len(slice))
	return (mean)
}

func MeanInt(slice []int) int {
	sumOfElements, _ := SumUp(slice)
	mean := sumOfElements / len(slice)
	return (mean)
}
