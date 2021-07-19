package main

import "fmt"

func main() {

	numChan := make(chan int, 100)
	resChan := make(chan map[int]int, 10)
	syncChan := make(chan<- bool, 8)

	//1.启动一个协程，将1-2000的数放到一个channel中，比如numChan
	go func(numChan chan int) {
		for i := 1; i <= 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}(numChan)

	//2.启动八个协程，从numChan取出数（比如n），并计算1+..+n的值，并存放到resChan
	for i := 0; i < 8; i++ {
		go func(numChan chan int, resChan chan map[int]int, syncChan chan<- bool) {
			for {
				n, ok := <-numChan
				if !ok {
					break
				}
				sum := 0
				for j := 1; j <= n; j++ {
					sum += j
				}
				numSum := make(map[int]int, 0)
				numSum[n] = sum
				resChan <- numSum
			}
			syncChan <- true
		}(numChan, resChan, syncChan)
	}

	go func() {
		for {
			if len(syncChan) == 8 {
				close(resChan)
				close(syncChan)
				break
			}
		}
	}()

	//3.最后八个协程协同完成工作后，再遍历resChan，并显示结果 [ 如 res[1]=1 .. res[10]=55 .. ]
	//4.注意，考虑resChan chan int 是否合适？
	for val := range resChan {
		for index, value := range val {
			fmt.Printf("res[%v]=%v\n", index, value)
		}
	}
}
