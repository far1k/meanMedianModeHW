package main

import (
	"fmt"
	"meanHomeWork/Mean"
	"meanHomeWork/Median"
	"meanHomeWork/Mode"
)

func main() {
	theSlice := []int{1, 2, 3, 3, 3, 4, 5}
	mean := Mean.MeanFloat(theSlice)
	median, _ := Median.Median(theSlice)
	mode, _ := Mode.Mode(theSlice)
	fmt.Printf("The numbers %d have following measures of central tendency:\nMean = %v\nMedian = %v\nMode = %v\n", theSlice, mean, median, mode)
}
