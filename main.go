package main

import (
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

func worker(wg *sync.WaitGroup, sudokuFiles <-chan string, id int) {
	defer wg.Done()
	for fileName := range sudokuFiles {
		time.Sleep(time.Millisecond)
		fmt.Printf("Id:%d Filename fetched: %s\n", id, fileName)
		s := new(Sudoku)
		s.Init(fileName)
		ss := new(SudokuSolver)

		if ss.Solve(s) == true {
			increaseSolved()
			fmt.Printf("Id:%d Sudoku is solved: %s\n", id, fileName)
		} else {
			increaseUnsolved()
			fmt.Printf("Id:%d Sudoku coudln't be solve: %s\n", id, fileName)
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

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, sudokuFiles, i)
	}

	for _, f := range files {
		fmt.Println("file added to channel: ", f.Name())
		sudokuFiles <- "./Dataset/" + f.Name()
	}
	close(sudokuFiles)
	wg.Wait()
	fmt.Printf("All sudoku files processed. Solved: %d Unsolved: %d\n", solved, unsolved)
}
