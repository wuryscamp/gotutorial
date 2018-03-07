package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	res1 := <-process(1, 3, &wg)
	res2 := <-process(2, 1, &wg)
	res3 := <-process(3, 4, &wg)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	wg.Wait()
}

func process(id int, delay time.Duration, wg *sync.WaitGroup) <-chan string {
	output := make(chan string)

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * delay)

		output <- fmt.Sprintf("Process %d done", id)
	}()

	return output
}
