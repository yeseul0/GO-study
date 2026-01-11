package main

import (
	"fmt"
	"strconv"
)

// 타입. 새자료형. 자료형 기본자료형에는 메소드 연결 못하니까 이렇게 함.
type MyInt int

func (i MyInt) Print() { //MyInt 타입에 대한 Print 메서드
	fmt.Println(i)
}

// type Rectangle struct {
// 	width, height int
// }

func (r Rectangle) Print() { //Rectangle 타입에 대한 Print 메서드
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print() //Print 메소드를 가진 모든 타입은 Printer 인터페이스 타입이 될 수 있음!!
}

func interfaceEx() {
	var i MyInt = 24
	i.Print()         //MyInt 타입의 메서드 직접 호출
	var p Printer = i //인터페이스 변수에 MyInt 타입 할당.
	//-> 얘는 줄여서 p := Printer(i) 이렇게도 쓸 수 있음..

	// 실제 MyInt에 더 많은 메소드가 있어도 인터페이스 p변수에 정의된 메소드만 호출 가능
	p.Print()

	r := Rectangle{10, 20}
	r.Print() //Rectangle 타입의 메서드 직접 호출
	p = r
	p.Print() //인터페이스 변수에 Rectangle 타입 할당 후 메서드 호출

	//-> 아! 메서드 집합만 같으면 같은 인터페이스 타입이 될 수 있구나!! duck typing

	fmt.Println(formatString(1))
	fmt.Println(formatString(3.14))
	fmt.Println(formatString("Hello"))
	fmt.Println(formatString(Person{"Alice", 30}))
}

// var x interface{ Print() }  // 이름 안 붙이고 바로 정의 (Print 함수를 갖고있는 인터페이스)
func f1(arg interface{}) { //빈 인터페이스: 모든 타입이 이 인터페이스가 될 수 있음
}

type Any interface{} //빈 인터페이스에 Any라는 이름 붙이기
func f2(arg Any) {
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	case int:
		i := arg.(int) //type assertion : interface{} -> int로 꺼내!
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string)
		return s
	default:
		return "Error"
	}
}
