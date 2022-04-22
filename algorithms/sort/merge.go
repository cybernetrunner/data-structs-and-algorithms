package sort

func MergeSort(items []int) []int {
	res := make(chan []int)

	go sort(items, res)

	sorted := <- res

	close(res)
	
	return sorted
}


func sort(items []int, result chan []int) {
	if len(items) < 2 {
		result <- items
		return
	}

	left  := make(chan []int)
	right := make(chan []int)
	middle := len(items)/2

	go sort(items[:middle], left)
	go sort(items[middle:], right)

	result <- merge(<-left, <-right)

	close(left)
	close(right)

	return
}

func merge(left []int, right []int) (result []int) {
	var i, j int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}
	
	for ; j < len(right); j++ {
		result = append(result, right[j])
	}
	
	return result
}
