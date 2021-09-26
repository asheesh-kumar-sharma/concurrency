package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(now))

}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("a... task 1")
}
func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("a... task 2 ")
}
func task3() {
	fmt.Println("a... task 3 ")
}
func task4() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("a... task 4")
}
