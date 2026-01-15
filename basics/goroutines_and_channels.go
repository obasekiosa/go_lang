package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(start, stop int, wg *sync.WaitGroup, done <-chan struct{}) {
	defer wg.Done()
	if start > stop {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()
		for index := start; ; index++ {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("Unbounded Routine: %d -- value %d\n", start, index)
			}
		}
	} else {
		for index := start; index < stop; index++ {
			fmt.Printf("Bounded Routine: %d -- value %d\n", start, index)
		}
	}
}

func runDoWorks() {
	length := 20
	batch_size := 4
	var wg sync.WaitGroup
	ch := make(chan struct{})

	for i := 0; i < length; i += batch_size {
		wg.Add(1)
		go doWork(i, i+batch_size, &wg, ch)
	}

	wg.Add(1)
	go doWork(length, 0, &wg, ch)

	var once sync.Once
	time.AfterFunc(5*time.Second, func() {
		once.Do(func() {
			close(ch)
		})
	})

	wg.Wait()
}
