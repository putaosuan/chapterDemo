package main

import "fmt"

func num(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}
func numhe(numChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		n, ok := <-numChan
		if !ok {
			break
		}
		num := 0
		map1 := make(map[int]int)
		for i := 0; i <= n; i++ {
			num += i
		}
		map1[n] = num
		resChan <- map1
	}
	exitChan <- true

}
func main() {
	numChan := make(chan int, 10)
	resChan := make(chan map[int]int, 10)
	exitChan := make(chan bool, 8)
	go num(numChan)
	for i := 0; i < 8; i++ {
		go numhe(numChan, resChan, exitChan)
	}
	go func() {
		for {
			if len(exitChan) == 8 {
				close(resChan)
				close(exitChan)
				break
			}
		}
	}()
	for val := range resChan {
		for index, value := range val {
			fmt.Printf("res[%v]=%v\n", index, value)
		}
	}

}
