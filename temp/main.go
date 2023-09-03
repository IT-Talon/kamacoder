package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println(r)
			}
		}()
		f1()
	}()
	fmt.Println("finish")
	time.Sleep(2)
}

func f1() {
	go f2()
	//panic("f1 panic")
	fmt.Println("f1")
}

func f2() {
	panic("f2 panic")
	fmt.Println("f2")

}
