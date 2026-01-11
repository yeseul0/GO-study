package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a int, b int, c chan int) {
	c <- a + b
}

// channel 기본 예제
func ex1() {

	//var c chan int
	//c = make(chan int)

	c := make(chan int) //int형 채널 생성

	go sum(1, 2, c) // sum 함수를 고루틴으로 실행, 채널을 매개변수로 넘겨줄 땐 꼭 go루틴 실행!
	n := <-c        //채널에서 값을 꺼내 n에 대입
	fmt.Println("Sum:", n)
}

// 동기 채널 예제
func ex2() {
	done := make(chan bool) //동기채널
	count := 3

	go func() { //익명 고루틴
		for i := 0; i < count; i++ {
			done <- true //done 채널에 true 보냄(값을 꺼낼 때까지 대기)
			fmt.Println("고루틴 : ", i)
			time.Sleep(1 * time.Second) //출력할 때 고루틴, 메인함수 순서로 예쁘게 출력하려고	^^
		}
	}() //고루틴 즉시 실행

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인함수 : ", i)
	}
}

// 비동기 채널 예제
func ex3() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) //버퍼크기 2인 비.동.기 채널
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true //done 채널에 true 보냄(버퍼가 꽉차면 대기)
			fmt.Println("고루틴 : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		time.Sleep(3 * time.Second)
		<-done //값을 꺼냄 (버퍼에 값이 없으면 대기)
		fmt.Println("메인함수 : ", i)
	}
	fmt.Scanln()
}

// close & range channel 예제
func ex4() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //채널 닫음
	}()

	for i := range c { //range 채널 이라는건, 채널이 닫힐 때까지를 의미
		fmt.Println(i)
	}
}

// producer-consumer 예제
func producer(c chan<- int) { //송신 전용 채널
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
	//<-c 는 컴파일 에러
}

func consumer(c <-chan int) { //수신 전용 채널
	for i := range c { //range로 채널에서 값을 꺼냄. (채널이 닫힐 때까지 계속)
		fmt.Println(i)
	}
	//c <- 200 는 컴파일 에러
}

func ex5() {
	c := make(chan int)
	go producer(c)
	go consumer(c)

	fmt.Scanln()
}

// 채널 파이프라인 예제
func num(a, b int) <-chan int { //int 수신 채널 리턴
	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out) //range 반복 끝내기 위해 닫음
	}()
	return out
}

func sum2(c <-chan int) <-chan int { //수신 채널 받아서 수신 채널 리턴
	out := make(chan int)
	go func() {
		r := 0
		for n := range c {
			r = r + n
		}
		out <- r
	}()
	return out
}
func main() {
	c := num(3, 5)
	out := sum2(c)
	fmt.Println(<-out)
}
