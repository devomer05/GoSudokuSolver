package sdk

import (
	"fmt"
	"io"
	"os"
)

// ROW constant
const (
	ROW         = 9
	NUMBERCOUNT = 9
	UNASSIGNED  = 0
)

// Sudoku struct
type Sudoku struct {
	data [ROW][ROW]int
}

// Init : Initialize sudoku with file
func (s *Sudoku) Init(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		var val int
		for i := 0; i < ROW; i++ {
			for j := 0; j < ROW; j++ {
				_, readErr := fmt.Fscanf(file, "%d", &val)
				if readErr == io.EOF {
					//fmt.Println("File read ended")
					break
				} else {
					s.Set(i, j, val)
				}
			}
			_, readErr := fmt.Fscanf(file, "\n", &val) // To skip new line char
			if readErr == io.EOF {
				//fmt.Println("File read ended")
				break
			}
		}
	}
}

func (s *Sudoku) usedInRow(row, val int) bool {
	for i := 0; i < ROW; i++ {
		if s.data[row][i] == val {
			return true
		}
	}
	return false
}

func (s *Sudoku) usedInCol(col, val int) bool {
	for i := 0; i < ROW; i++ {
		if s.data[i][col] == val {
			return true
		}
	}
	return false
}

func (s *Sudoku) usedInBox(x, y, val int) bool {
	boxStartRow := x - x%3
	boxStartCol := y - y%3

	for row := 0; row < ROW/3; row++ {
		for col := 0; col < ROW/3; col++ {
			if s.data[boxStartRow+row][col+boxStartCol] == val {
				return true
			}
		}
	}
	return false
}

// Set a value to given cell
func (s *Sudoku) Set(x, y, val int) {
	s.data[x][y] = val
}

// Get the value of cell
func (s *Sudoku) Get(x, y int) int {
	return s.data[x][y]
}

// IsSafe checks cell
func (s *Sudoku) IsSafe(x, y, val int) bool {
	return !s.usedInRow(x, val) && !s.usedInCol(y, val) && !s.usedInBox(x, y, val)
}

// GetFirstUnassigned cell position
func (s *Sudoku) GetFirstUnassigned() (x, y int, found bool) {
	x, y, found = 0, 0, false
	for i := 0; i < ROW; i++ {
		for j := 0; j < ROW; j++ {
			if s.data[i][j] == UNASSIGNED {
				x = i
				y = j
				found = true
				return
			}
		}
	}
	return
}

// Print prints the sudoku
func (s *Sudoku) Print() {
	fmt.Println("\t  | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |")
	fmt.Println("\t  -------------------------------------")
	for i := 0; i < ROW; i++ {
		fmt.Print("\t", i+1, " | ")
		for j := 0; j < ROW; j++ {
			fmt.Printf("%d ", s.data[i][j])
			if (j+1)%3 == 0 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println()
		if (i+1)%3 == 0 {
			fmt.Println("\t-------------------------------------")
		} else {
			fmt.Println("\t                                     ")
		}
	}
}

// CreateSudoku creates a new sudoku and returns its pointer
func CreateSudoku(fileName string) *Sudoku {
	s := new(Sudoku)
	s.Init(fileName)
	return s
}
