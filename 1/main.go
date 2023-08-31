package main

import (
	"fmt"
	"time"
)

func main() {
	//for {
	//	var a, b int
	//	_, err := fmt.Scan(&a, &b)
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(a + b)
	//}
	var c chan int
	i, ok := <-c
	fmt.Println(i)
	fmt.Println(ok)
}

func insert(c chan int) {
	for i := 0; i < 4; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	close(c)

}
