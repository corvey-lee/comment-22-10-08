package main

import (
	"log"
	"testing"
	"time"
)

func TestAsyncTiming(t *testing.T) {
	fileListCount := 100

	start := time.Now()
	Async(fileListCount)
	log.Println("비동기 걸린 시간: ", time.Now().Sub(start))
}

func TestSyncTiming(t *testing.T) {
	fileListCount := 100

	start := time.Now()
	Sync(fileListCount)
	log.Println("동기 걸린 시간: ", time.Now().Sub(start))
}
