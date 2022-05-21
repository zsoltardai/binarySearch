package binarySearch

func BinarySearch(sequence []int, item int) int {

	var starting_point int = 0
	var ending_point int = len(sequence) - 1

	for ok := true; ok; ok = starting_point <= ending_point {

		var midpoint int = starting_point + (ending_point-starting_point)/2

		if sequence[midpoint] == item {
			return midpoint
		} else if sequence[midpoint] < item {
			starting_point = midpoint + 1
		} else {
			ending_point = midpoint - 1
		}
	}

	return -1
}