package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 뮤텍스 예제 (Lock, Unlock, Gosched)
func ex1() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.GOMAXPROCS(1)) //사용 가능한 모든 CPU 코어 사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 100; i++ {
			mutex.Lock() //뮤텍스 잠금
			data = append(data, 1)
			mutex.Unlock() //뮤텍스 잠금 해제

			runtime.Gosched() //다른 고루틴이 CPU 사용할 수 있도록 실행 양보
		}
	}()

	//같은 함수 다른 고루틴
	go func() {
		for i := 0; i < 100; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))
}

// sync.RWMutex 예제
func ex2() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 코어 사용

	var data int = 0
	var rwMutex = new(sync.RWMutex)

	go func() { // 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.Lock()                    // 쓰기 잠금
			data += 1                         // data에 값 쓰기
			fmt.Println("write   : ", data)   // data 값을 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
			rwMutex.Unlock()                  // 쓰기 잠금 해제
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()                // 읽기 잠금
			fmt.Println("read 1 : ", data) // data 값을 출력(읽기)
			time.Sleep(1 * time.Second)    // 1초 대기
			rwMutex.RUnlock()              // 읽기 잠금 해제
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()                // 읽기 잠금
			fmt.Println("read 2 : ", data) // data 값을 출력(읽기)
			time.Sleep(2 * time.Second)    // 2초 대기
			rwMutex.RUnlock()              // 읽기 잠금 해제
		}
	}()

	time.Sleep(10 * time.Second) // 10초
}

//읽을땐 여러 고루틴이 동시에 읽을 수 있지만, 쓸땐 오직 하나의 고루틴만 쓸 수 있음
//쓸 땐 읽기도 금지. 읽을땐 쓰기만 금지

// sync.Cond & Signal
func ex3() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 c 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true    // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait()                   // 조건 변수 대기
			fmt.Println("wait end : ", n) //signal 받고, mutex 잡으면 실행됨. (mutex를 잡는다는건 signal이 lock하고 있던 mutex를 unlock했다는 것)
			mutex.Unlock()                // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄
	}
	// 고루틴 3개가 모두 실행돼서 3개의 값이 꺼내질때까지 대기

	for i := 0; i < 3; i++ {
		mutex.Lock() // 뮤텍스 잠금, cond.Signal() 보호 시작
		fmt.Println("signal : ", i)
		cond.Signal()  // 대기하고 있는 고루틴을 하나씩 깨움
		mutex.Unlock() // 뮤텍스 잠금 해제, cond.Signal() 보고 종료
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Scanln()
}

/*
	cond.Wait()은 호출한 고루틴을 대기 상태로 만들고, 뮤텍스를 unlock함.
	고루틴1 unlock -> 고루틴2 lock -> 고루틴2 unlock -> 고루틴3 lock -> 고루틴3 unlock
	cond.Signal()은 대기 상태인 고루틴 하나를 깨우고, 그 고루틴은 다시 뮤텍스를 lock함.
	cond.Signal()이 고루틴1 깨움 -> 고루틴1 lock -> 고루틴1 unlock
	-> cond.Signal()이 고루틴2 깨움 -> 고루틴2 lock -> 고루틴2 unlock
	-> cond.Signal()이 고루틴3 깨움 -> 고루틴3 lock -> 고루틴3 unlock

	출력:
	wait begin :  2
	wait begin :  1
	wait begin :  0
	signal :  0
	signal :  1
	signal :  2
	wait end :  2
	wait end :  1
	wait end :  0

	signal, wait end가 번갈아 나오지 않고 signal이 먼저 다 출력되는 이유.
	메인 for문이 더 빨라서 깨어난 고루틴이 뮤텍스 잡기 전에 다음 for문이 mutex를 잡아버리기 때문!!

	따라서 time.Sleep() 추가하면 signal, wait end가 번갈아 출력됨
*/

// sync.Cond & Broadcast
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true    // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	mutex.Lock() // 뮤텍스 잠금, cond.Broadcast() 보호 시작
	fmt.Println("broadcast")
	cond.Broadcast() // 대기하고 있는 모.든 고루틴을 깨움
	mutex.Unlock()   // 뮤텍스 잠금 해제, cond.Signal() 보고 종료

	fmt.Scanln()
}
