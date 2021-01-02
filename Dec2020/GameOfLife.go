package Dec2020

func GameOfLife(board [][]int) {
	const live = 2
	const dead = 3

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			cnt := 0
			if i > 0 && j > 0 {
				cnt += getPrevValue(board[i-1][j-1])
			}
			if i > 0 {
				cnt += getPrevValue(board[i-1][j])
				if j < len(board[0])-1 {
					cnt += getPrevValue(board[i-1][j+1])
				}
			}
			if j > 0 {
				cnt += getPrevValue(board[i][j-1])
				if i < len(board)-1 {
					cnt += getPrevValue(board[i+1][j-1])
				}
			}
			if i < len(board)-1 && j < len(board[0])-1 {
				cnt += getPrevValue(board[i+1][j+1])
			}
			if j < len(board[0])-1 {
				cnt += getPrevValue(board[i][j+1])
			}
			if i < len(board)-1 {
				cnt += getPrevValue(board[i+1][j])
			}
			if cnt < 2 {
				if board[i][j] != 0 {
					board[i][j] = dead
				}
			} else if cnt > 3 {
				if board[i][j] != 0 {
					board[i][j] = dead
				}
			} else if cnt == 3 {
				if board[i][j] != 1 {
					board[i][j] = live
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == dead {
				board[i][j] = 0
			} else if board[i][j] == live {
				board[i][j] = 1
			}
		}
	}
}

func getPrevValue(a int) int {
	if a == 2 {
		return 0
	} else if a == 3 {
		return 1
	}
	return a
}
