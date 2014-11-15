package board

type Cell bool

const (
	ALIVE Cell = true
	DIE Cell = false
)

func AliveNeighbors(board [][]Cell, row, column int) int {
	count := 0
	count += AliveTop(board, row, column)
	count += AliveTopRight(board, row, column)
	count += AliveRight(board, row, column)
	count += AliveBottomRight(board, row, column)
	count += AliveBottom(board,row, column)
	count += AliveBottomLeft(board, row, column)
	count += AliveLeft(board, row, column)
	count += AliveTopLeft(board, row, column)
	return count
}

func AliveTop(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row-1, column) && board[row-1][column] == ALIVE) {
		return 1
	}
	return 0
}

func AliveTopRight(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row-1, column+1) && board[row-1][column+1] == ALIVE) {
		return 1
	}
	return 0
}

func AliveRight(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row, column+1) && board[row][column+1] == ALIVE) {
		return 1
	}
	return 0
}

func AliveBottomRight(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row+1, column+1) && board[row+1][column+1] == ALIVE) {
		return 1
	}
	return 0
}

func AliveBottom(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row+1, column) && board[row+1][column] == ALIVE) {
		return 1
	}
	return 0
}

func AliveBottomLeft(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row+1, column-1) && board[row+1][column-1] == ALIVE) {
		return 1
	}
	return 0
}

func AliveLeft(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row, column-1) && board[row][column-1] == ALIVE) {
		return 1
	}
	return 0
}

func AliveTopLeft(board [][]Cell, row, column  int) int {
	if (IsInBoard(board, row-1, column-1) && board[row-1][column-1] == ALIVE) {
		return 1
	}
	return 0
}


func IsInBoard(board [][]Cell, row, column int) bool {
	maxRow := len(board)
	maxColumn := len(board[0])
	return (row >= 0 && row < maxRow && column >= 0 && column < maxColumn)
}

func NextGenerationBoard(board [][]Cell) (nextBoard [][]Cell) {
	nextBoard = CopyBoard(board)
	maxRow := len(board)
	maxColumn := len(board[0])

	for row := 0; row < maxRow; row++ {
		for column := 0; column < maxColumn; column++ {
			neighbors := AliveNeighbors(board, row, column)
			nextBoard[row][column] = NextGenerationCell(board[row][column], neighbors)
		}
	}
	return 
}

func NextGenerationCell(cell Cell, neighbors int) Cell{
	if cell == ALIVE {
		if neighbors == 2 || neighbors == 3 {
			return ALIVE
		}
	}

	if neighbors == 3 {
		return ALIVE
	}

	return DIE
}

func CopyBoard(src [][]Cell) (dst [][]Cell) {
	maxRow := len(src)
	dst = make([][]Cell, maxRow)
	for i, row := range src {
		maxColumn := len(row)
		dst[i] = make([]Cell, maxColumn)
		for j, col := range row {
			dst[i][j] = col
		}
	}
	return
}
