package sdk

// SudokuSolver struct
type SudokuSolver struct {
	foundCellCount int
}

// Solve given sudoku
func (ss *SudokuSolver) Solve(s *Sudoku) bool {

	return ss.solve(s)
}

func (ss *SudokuSolver) solve(s *Sudoku) bool {
	x, y, ok := s.GetFirstUnassigned()
	if ok == false {
		return true
	}

	for i := 1; i <= NUMBERCOUNT; i++ {
		if s.IsSafe(x, y, i) {
			s.Set(x, y, i)

			if ss.solve(s) {
				return true
			}
			s.Set(x, y, UNASSIGNED)
		}
	}
	return false
}

// CreateSolver creates a sudoku solver
func CreateSolver() *SudokuSolver {
	ss := new(SudokuSolver)
	return ss
}
