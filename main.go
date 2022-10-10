package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	fileListCount := 100

	// fmt.Println(Sync(fileListCount))
	fmt.Println("Romeo 단어의 총 갯수: ", Async(fileListCount))
}

func Async(fileListCount int) int {
	total := 0
	ch := make(chan int, fileListCount)
	wg := sync.WaitGroup{}
	for i := 0; i < fileListCount; i++ {
		path := "./sample.txt"
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- CountingStrings(path)
		}()
	}
	wg.Wait()
	close(ch)

	for i := range ch {
		total += i
	}

	return total
}

func Sync(fileListCount int) int {
	total := 0

	for i := 0; i < fileListCount; i++ {
		path := "./sample.txt"
		total += CountingStrings(path)
	}

	return total
}

func CountingStrings(path string) int {
	sum := 0

	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Failed to read file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Romeo") {
			sum += 1
		}
		line++
	}

	return sum
}
