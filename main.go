package main

import (
	"SudokuSolverGo/sdk"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

var mutexSolved = &sync.Mutex{}
var mutexUnsolved = &sync.Mutex{}

var solved int
var unsolved int

func increaseSolved() {
	mutexSolved.Lock()
	solved++
	mutexSolved.Unlock()
}

func increaseUnsolved() {
	mutexUnsolved.Lock()
	unsolved++
	mutexUnsolved.Unlock()
}

func worker(wg *sync.WaitGroup, sudokuFiles <-chan string, id int, ss *sdk.SudokuSolver) {
	defer func(wg *sync.WaitGroup) {
		fmt.Printf("Goroutine: %d Ended\n", id)
		wg.Done()
	}(wg)

	fmt.Printf("Goroutine: %d Started\n", id)
	for fileName := range sudokuFiles {
		time.Sleep(time.Millisecond)
		fmt.Printf("Goroutine: %d Filename fetched: %s\n", id, fileName)
		s := sdk.CreateSudoku(fileName)

		if ss.Solve(s) == true {
			increaseSolved()
			fmt.Printf("Goroutine: %d Sudoku is solved: %s\n", id, fileName)
		} else {
			increaseUnsolved()
			fmt.Printf("Goroutine: %d Sudoku coudln't be solve: %s\n", id, fileName)
		}
	}
}

func main() {

	var wg sync.WaitGroup
	sudokuFiles := make(chan string, 20)

	files, err := ioutil.ReadDir("./Dataset/")
	if err != nil {
		log.Fatal(err)
	}

	ss := sdk.CreateSolver()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, sudokuFiles, i, ss)
	}

	for _, f := range files {
		fmt.Println("file added to channel: ", f.Name())
		sudokuFiles <- "./Dataset/" + f.Name()
	}
	close(sudokuFiles)
	wg.Wait()
	fmt.Printf("All sudoku files processed. Solved: %d Unsolved: %d\n", solved, unsolved)
}
