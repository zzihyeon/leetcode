package Dec2021

func CanFormArray(arr []int, pieces [][]int) bool {
	inOrderMap := make(map[int]int) //[array value]piece index
	for i, piece := range pieces {
		for _, ele := range piece {
			inOrderMap[ele] = i + 1
		}
	}
	for i := 0; i < len(arr); i++ {
		piecesIdx := inOrderMap[arr[i]]
		if piecesIdx == 0 {
			return false
		}
		for j := i; j < i+len(pieces[piecesIdx-1]); j++ {
			if inOrderMap[arr[j]] != piecesIdx {
				return false
			} else if pieces[piecesIdx-1][j-i] != arr[j] {
				return false
			}
		}
		i = i + len(pieces[piecesIdx-1]) - 1
	}
	return true
}
