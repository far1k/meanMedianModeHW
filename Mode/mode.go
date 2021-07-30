package Mode

import "errors"

func Mode(slice []int) ([]int, error) {
	//       Account for empty slice case
	if len(slice) == 0 {
		return nil, errors.New("Please enter a non-empty slice")
	}
	// 1) CREATE A FREQUENCY DISTRIBUTION MAP
	//       Initialising a Map
	Map := map[int]int{
		slice[0]: 0,
	}
	//       Declaring unique values in slice as KEYS in map, and frequency of appearence in slice as VALUES in map
	for i := range slice {
		if val, ok := Map[slice[i]]; ok {
			Map[slice[i]] = val + 1
		} else {
			Map[slice[i]] = 1
		}
	}
	// 2) FIND MOST COMMONLY APPEARING VALUE(S)
	//       Iterate through Map to find out which value(s) appear most often and write these values in an output slice
	maxCounter := 0
	sliceOfModes := []int{}
	for key, val := range Map {
		if val > maxCounter {
			maxCounter = val
			sliceOfModes = []int{}
			sliceOfModes = append(sliceOfModes, key)
		} else if val == maxCounter {
			sliceOfModes = append(sliceOfModes, key)
		}
	}
	//       Return values & account for "equally often" case
	if len(sliceOfModes) == len(Map) {
		return nil, errors.New("All values appear equally often, therefore the dataset has NO MODE")
	} else {
		return (sliceOfModes), nil
	}
}
