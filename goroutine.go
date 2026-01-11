package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main() {
	runtime.GOMAXPROCS(1) //CPU 코어 1개만 사용하도록!
	fmt.Println("system total CPU cores:", runtime.NumCPU())
	fmt.Println("runtime cpu cores used: ", runtime.GOMAXPROCS(0))

	s := "Hello, world"
	for i := 0; i < 20; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i) //선언 하자마자 i를 n으로 복사 넘기면서 바로 호출 고루틴으로 익명함수 실행
	} //고루틴 순서가 보장되지는 않음 -> 그건 동기채널 써야함!
	/*
		만약에!!
		for i := 0; i < 20; i++ {
			go func() {
				fmt.Println(s, i)
			}() i를 변수로 안 넘겨 받고 직접 출력하게되면 다 20임...
			 왜냐면 고루틴으로 실행한 클로저는 for문이 다 돌고 i가 20이 된 상태이기 때문
		}
		-> 고루틴으로 실행하는 클로저에 반복문에 의해 바뀌는 변수는 꼭 매개변수로 넘기자!
	*/
	fmt.Scanln() //메인 고루틴이 종료되지 않도록 대기(enter로 종료)
}
