package main //이 file은 main pck에 속함

import "fmt" //fmt=format output 라이브러리

func main() {
	fmt.Println("Hello World")
	FizzBuzz(100)
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz") //3,5 공배수
		case i%3 == 0:
			fmt.Println("Fizz") //3의 배수
		case i%5 == 0:
			fmt.Println("Buzz") //5의 배수
		default:
			fmt.Println(i) //그 외 숫자
		}
	}
}
