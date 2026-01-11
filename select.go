package main

import (
	"fmt"
	"time"
)

func ex1() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10 //채널 c1에 10보냄
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "hello world" //채널 c2에 "hello world" 보냄
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			// case <-c1 으로 써도 됨
			case i := <-c1: //채널 c1에 값이 들어왔다면 꺼내서 i에 대입
				fmt.Println("cl :", i)
			case s := <-c2: //채널 c2에 값이 들어왔다면 꺼내서 s에 대입
				fmt.Println("c2 :", s)
			case <-time.After(50 * time.Millisecond): //50ms 후 현재 시간이 담긴 채널 리턴됨
				fmt.Println("timeout 50ms")
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func ex2() {
	c1 := make(chan int)    // int형 채널 생성
	c2 := make(chan string) // string 채널 생성

	go func() {
		for {
			i := <-c1                          // 채널 c1에서 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1 :", i)             // i 값을 출력
			time.Sleep(100 * time.Millisecond) // 100 밀리초 대기
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"              // 채널 c2에 Hello, world!를 보냄
			time.Sleep(500 * time.Millisecond) // 500 밀리초 대기
		}
	}()

	go func() {
		for { // 무한 루프
			select {
			case c1 <- 10: // 매번 채널 c1에 10을 보냄
			case s := <-c2: // c2에 값이 들어왔을 때는 값을 꺼낸 뒤 s에 대입
				fmt.Println("c2 :", s) // s 값을 출력
			}
		} //매번 채널 c1에 10을 보내면서, 채널 c2에 값이 들어오면 그 값을 출력
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}

func main() {
	c1 := make(chan int) // int형 채널 생성

	go func() {
		for {
			i := <-c1                          // 채널 c1에서 값을 꺼낸 뒤 i에 대입
			fmt.Println("81줄 c1 :", i)         // i 값을 출력
			time.Sleep(100 * time.Millisecond) // 100 밀리초 대기
		}
	}()

	go func() {
		for {
			c1 <- 20                           // 채널 c1에 20을 보냄
			time.Sleep(500 * time.Millisecond) // 100 밀리초 대기
		}
	}()

	go func() {
		for { // 무한 루프
			select { // 채널 c1 한 개로 값을 보내거나 받음
			case c1 <- 10: // 매번 채널 c1에 10을 보냄 -> 81줄에서 꺼내줌
			case i := <-c1: // 88줄에서넣은거 여기서 꺼냄
				fmt.Println("98 줄 c1 :", i) // i 값을 출력
			}
		}
	}()

	// select 분기분에서 채널에 넣으면, 다른 쪽에서 꺼내고,
	// 다른 쪽에서 넣으면 select 분기문에서 꺼내는 식으로 동작됨
	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행

}
