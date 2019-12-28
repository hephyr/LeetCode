func isValidSudoku(board [][]byte) bool {
    seen := make(map[string]int)
    for i:= 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            c := board[i][j]
            if c == '.' {
                continue
            }
            row := "row " + strconv.Itoa(i) + " has " + string(c)
            col := "col " + strconv.Itoa(j) + " has " + string(c)
            box := "box " + strconv.Itoa(i/3) + strconv.Itoa(j/3) + " has " + string(c)
            if isKeyInMap(seen, row) || isKeyInMap(seen, col) || isKeyInMap(seen, box) {
                return false
            }
        }
    }
    return true
}

func isKeyInMap(m map[string]int, key string) bool {
    _, ok := m[key]
    if !ok {
        m[key] = 1
    }
    return ok
}