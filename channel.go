package main

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b
}

func main() {
	c := make(chan int) //int형 채널 생성
	go sum(1, 2, c)     // sum 함수를 고루틴으로 실행, 채널을 매개변수로 넘겨줌
	n := <-c            //채널에서 값을 꺼내 n에 대입
	fmt.Println("Sum:", n)
}
