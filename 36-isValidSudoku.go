package leetcode

//func isValidSudoku(board [][]byte) bool {
//	isValid := func(board [][]byte) bool {
//		m := map[byte]bool{}
//		for i := 0; i < 3; i++ {
//			for j := 0; j < 3; j++ {
//				if board[i][j] == '.' {
//					continue
//				}
//				if _, ok := m[board[i][j]]; ok {
//					return false
//				}
//				m[board[i][j]] = true
//			}
//		}
//		return true
//	}
//	for i := 0; i < 9; i++ {
//		m := map[byte]bool{}
//		mm := map[byte]bool{}
//		for j := 0; j < 9; j++ {
//			if _, ok := m[board[i][j]]; ok {
//				return false
//			}
//			if board[i][j] != '.'{
//				m[board[i][j]] = true
//			}
//			if _, ok := mm[board[j][i]]; ok && board[j][i] != '.'{
//				return false
//			}
//			if board[j][i] != '.'{
//				mm[board[j][i]] = true
//			}
//		}
//	}
//	for i := 0; i < 9; i += 3 {
//		for j := 0; j < 9; j += 3 {
//			var arr [][]byte
//			arr = append(arr, board[i][j:j+3])
//			arr = append(arr, board[i+1][j:j+3])
//			arr = append(arr, board[i+2][j:j+3])
//			if !isValid(arr) {
//				return false
//			}
//		}
//	}
//	return true
//}


func isValidSudoku(board [][]byte) bool {
	var rows,colums  [9][9]int
	var sub  [3][3][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.'{
				continue
			}
			idx := board[i][j]-'1'
			rows[i][idx]++
			colums[j][idx]++
			sub[i/3][j/3][idx]++
			if rows[i][idx] >1 || colums[j][idx]>1 || sub[i/3][j/3][idx] >1{
				return false
			}
		}
	}
	return true
}
