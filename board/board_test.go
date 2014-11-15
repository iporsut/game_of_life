package board

import (
	"testing"
	"reflect"
)

func TestCountAliveNeighborsSingleCellIsZero(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE },
	}

	count := AliveNeighbors(board, 0, 0)

	if count != 0 {
		t.Errorf("Expect 0 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsTopLeftCellIsThree(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE },
		[]Cell { ALIVE, ALIVE },
	}

	count := AliveNeighbors(board, 0, 0)

	if count != 3 {
		t.Errorf("Expect 3 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsTopLeftCellForFourFourGridIsThree(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 0, 0)

	if count != 3 {
		t.Errorf("Expect 3 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsTopRightCellIsThree(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 0, 3)

	if count != 3 {
		t.Errorf("Expect 3 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsBottomRightCellIsThree(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 3, 3)

	if count != 3 {
		t.Errorf("Expect 3 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsBottomLeftCellIsThree(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 3, 0)

	if count != 3 {
		t.Errorf("Expect 3 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsTopCellIsFive(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 0, 1)

	if count != 5 {
		t.Errorf("Expect 5 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsRightCellIsFive(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 1, 3)

	if count != 5 {
		t.Errorf("Expect 5 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsBottomCellIsFive(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 3, 1)

	if count != 5 {
		t.Errorf("Expect 5 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsLeftCellIsFive(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 1, 0)

	if count != 5 {
		t.Errorf("Expect 5 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsOneOneCellIsEight(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 1, 1)

	if count != 8 {
		t.Errorf("Expect 8 alive neighbors but %v", count)
	}
}

func TestCountAliveNeighborsTwoTwoCellIsEight(t *testing.T) {
	board := [][]Cell {
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
		[]Cell { ALIVE, ALIVE, ALIVE ,ALIVE},
	}

	count := AliveNeighbors(board, 2, 2)

	if count != 8 {
		t.Errorf("Expect 8 alive neighbors but %v", count)
	}
}

func TestCopyBoard(t *testing.T) {
	board := [][]Cell {
		[]Cell { DIE, DIE, DIE},
		[]Cell { DIE, ALIVE, DIE},
		[]Cell { DIE, DIE, DIE},
	}

	var newBoard [][]Cell = CopyBoard(board)

	if !reflect.DeepEqual(newBoard, board) {
		t.Errorf("Expect %v but %v", board, newBoard)
	}
}

func TestNextGenerationBoardFirstAliveRule(t *testing.T) {
	board := [][]Cell {
		[]Cell { DIE, DIE, DIE},
		[]Cell { DIE, ALIVE, DIE},
		[]Cell { DIE, DIE, DIE},
	}

	expect := [][]Cell {
		[]Cell { DIE, DIE, DIE},
		[]Cell { DIE, DIE, DIE},
		[]Cell { DIE, DIE, DIE},
	}

	nextBoard := NextGenerationBoard(board)

	if !reflect.DeepEqual(nextBoard, expect) {
		t.Errorf("Expect nextboard %v but %v", expect, nextBoard)
	}
}

func TestNextGenerationBoardStableAlive(t *testing.T) {
	board := [][]Cell {
		[]Cell { DIE, DIE, DIE, DIE},
		[]Cell { DIE, ALIVE, ALIVE, DIE},
		[]Cell { DIE, ALIVE, ALIVE, DIE},
		[]Cell { DIE, DIE, DIE, DIE},
	}

	expect := [][]Cell {
		[]Cell { DIE, DIE, DIE, DIE},
		[]Cell { DIE, ALIVE, ALIVE, DIE},
		[]Cell { DIE, ALIVE, ALIVE, DIE},
		[]Cell { DIE, DIE, DIE, DIE},
	}

	nextBoard := NextGenerationBoard(board)

	if !reflect.DeepEqual(nextBoard, expect) {
		t.Errorf("Expect nextboard %v but %v", expect, nextBoard)
	}
}

func TestNextGenerationBoardOscillators(t *testing.T) {
	board := [][]Cell {
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, DIE, ALIVE, DIE, DIE},
		[]Cell {DIE, DIE, ALIVE, DIE, DIE},
		[]Cell {DIE, DIE, ALIVE, DIE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
	}

	expect := [][]Cell {
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, ALIVE, ALIVE, ALIVE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
	}

	nextBoard := NextGenerationBoard(board)

	if !reflect.DeepEqual(nextBoard, expect) {
		t.Errorf("Expect nextboard %v but %v", expect, nextBoard)
	}
}

func TestNextGenerationBoardOscillatorsBack(t *testing.T) {
	board := [][]Cell {
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, ALIVE, ALIVE, ALIVE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
	}

	expect := [][]Cell {
		[]Cell {DIE, DIE, DIE, DIE, DIE},
		[]Cell {DIE, DIE, ALIVE, DIE, DIE},
		[]Cell {DIE, DIE, ALIVE, DIE, DIE},
		[]Cell {DIE, DIE, ALIVE, DIE, DIE},
		[]Cell {DIE, DIE, DIE, DIE, DIE},
	}

	nextBoard := NextGenerationBoard(board)

	if !reflect.DeepEqual(nextBoard, expect) {
		t.Errorf("Expect nextboard %v but %v", expect, nextBoard)
	}
}
