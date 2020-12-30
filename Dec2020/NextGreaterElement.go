package Dec2020

func NextGreaterElement(n int) int {
	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)
	given := n
	numToArray := []int{}
	for ; n > 9; n = n / 10 {
		digit := n % 10
		numToArray = append(numToArray, digit)
	}
	numToArray = append(numToArray, n)
	nextArray, idx := makeNGE(numToArray, given)
	nextArray = restSort(nextArray, idx)
	result := arrayToInt(nextArray)
	if given == result || given > result || result > int(MaxInt) {
		return -1
	}
	return result
}

func makeNGE(numToArray []int, given int) ([]int, int) {
	prev := numToArray[0]
	history := map[int][]int{}
	ptr := 0
	for index, value := range numToArray {
		if prev > value {
			ptr = index
			minIdx := findBiggerValueIndex(history, value)
			temp := numToArray[minIdx]
			numToArray[minIdx] = numToArray[index]
			numToArray[index] = temp
			break
		}
		history[value] = append(history[value], index)
		prev = value
	}
	return numToArray, ptr
}

func restSort(numToArray []int, idx int) []int {
	for i := 0; i < idx-1; i++ {
		for j := i + 1; j < idx; j++ {
			if numToArray[i] < numToArray[j] {
				temp := numToArray[i]
				numToArray[i] = numToArray[j]
				numToArray[j] = temp
			}
		}
	}
	return numToArray
}

func arrayToInt(array []int) int {
	result := 0
	target := 1
	for _, ele := range array {
		result = result + ele*target
		target = target * 10
	}
	return result
}

func findBiggerValueIndex(history map[int][]int, value int) int {
	i := value + 1
	for ; i < 10; i++ {
		if len(history[i]) > 0 {
			break
		}
	}
	return history[i][0]
}
