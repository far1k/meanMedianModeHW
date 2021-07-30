package Median

import (
	"errors"
	"meanHomeWork/Mean"
)

func Median(slice []int) (float64, error) {
	//       Account for empty slice case
	if len(slice) == 0 {
		return 0, errors.New("Please enter a non-empty slice")
	}
	// 1) SORT ALL THE VALUES IN ASCENDING ORDER
	sortedSlice := []int{slice[0]}
	//      Iterate through slice and compare each value to values in the sortedList, if value is greater than
	//      the previous one but lower that the next one -> insert it there
	for i := 1; i < len(slice); i++ {
	Loop:
		for j := range sortedSlice {
			if slice[i] <= sortedSlice[j] {
				sortedSlice = append(sortedSlice, 0)
				copy(sortedSlice[(j+1):], sortedSlice[j:])
				sortedSlice[j] = slice[i]
				break Loop
			} else if slice[i] > sortedSlice[len(sortedSlice)-1] {
				sortedSlice = append(sortedSlice, slice[i])
				break Loop
			}
		}
	}
	// 2) FIND AND RETURN THE MIDDLE NUMBER(S)
	slice = sortedSlice
	if len(slice)%2 != 0 {
		return (float64(slice[len(slice)/2])), nil
	//      If there is even number of elements in input slice, then the middle will be two numbers,
	//      in which case we take the average
	} else {
		return (Mean.MeanFloat([]int{slice[len(slice)/2-1], slice[len(slice)/2]})), nil
	}
}
