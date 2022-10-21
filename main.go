package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fileListCount := 100

	// fmt.Println("(동기) Romeo 단어의 총 갯수: ", Sync(fileListCount))
	fmt.Println("(비동기) Romeo 단어의 총 갯수: ", Async(fileListCount))
}

func Async(fileListCount int) int {
	total := 0
	ch := make(chan int, fileListCount)
	wg := sync.WaitGroup{}
	temp_path := "./temp"

	// temp 폴더 생성
	if _, err := os.Stat(temp_path); os.IsNotExist(err) {
		err := os.Mkdir(temp_path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// temp 폴더에 sampleN.txt 파일 생성하고 삭제
	for i := 0; i < fileListCount; i++ {
		path := temp_path + "/sample" + strconv.Itoa(i) + ".txt"
		cmd := exec.Command("cp", "./sample.txt", path)
		_, err := cmd.Output()
		if err != nil {
			log.Println(err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- countingStrings(path)
			os.Remove(path)
		}()
	}
	wg.Wait()
	close(ch)

	for i := range ch {
		total += i
	}

	// select {
	// case i := <-ch:
	// 	total += i
	// }

	return total
}

func Sync(fileListCount int) int {
	total := 0
	temp_path := "./temp"

	// temp 폴더 생성
	if _, err := os.Stat(temp_path); os.IsNotExist(err) {
		err := os.Mkdir(temp_path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// temp 폴더에 sampleN.txt 파일 생성하고 삭제
	for i := 0; i < fileListCount; i++ {
		path := temp_path + "/sample" + strconv.Itoa(i) + ".txt"
		cmd := exec.Command("cp", "./sample.txt", path)
		_, err := cmd.Output()
		if err != nil {
			log.Println(err)
		}

		total += countingStrings(path)
		os.Remove(path)
	}

	return total
}

func countingStrings(path string) int {
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
