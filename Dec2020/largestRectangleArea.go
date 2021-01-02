package Dec2020

func LargestRectangleArea(heights []int) int {
	already := make(map[int]bool)
	maxArea := 0
	for _, data := range heights {
		if already[data] != true {
			already[data] = true
			res := maxhor(heights, data)
			if res*data > maxArea {
				maxArea = res * data
			}
		}
	}
	return maxArea
}

func maxhor(heights []int, data int) int {
	max := 0
	cnt := 0
	for _, ele := range heights {
		if data > ele {
			if max < cnt {
				max = cnt
			}
			cnt = 0
		} else {
			cnt++
		}
	}
	if max < cnt {
		max = cnt
	}
	return max
}
