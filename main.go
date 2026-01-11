package main //이 file은 main pck에 속함

import "fmt" //fmt=format output 라이브러리

func main() {
	fmt.Println("Hello World")
	// FizzBuzz(10)
	// n_BottlesOfBeer(5)

	pointerEx()
	structPointerEx()
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ { // :=는 var과 타입 압축
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

func n_BottlesOfBeer(n int) {
	for i := n; i >= 0; i-- {
		switch {
		case i > 1:
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", i, i)
			s := "bottles"
			if i-1 == 1 {
				s = "bottle"
			}
			fmt.Printf("Take one down and pass it around, %d %s of beer on the wall.\n\n", i-1, s)
		case i == 1:
			fmt.Printf("1 bottle of beer on the wall, 1 bottle of beer.\n")
			fmt.Printf("Take one down and pass it around, no more bottles of beer on the wall.\n\n")
		default:
			fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\n")
			fmt.Printf("Go to the store and buy some more, %d bottles of beer on the wall.\n", n)
		}
	}
}
