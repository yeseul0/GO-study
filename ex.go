package main

import "fmt"

func pointerEx() {
	//int형 변수를 가리키는 포인터 변수 선언
	var numPtr *int = new(int) //new로 메모리 할당 안해주면 nil포인터임!!!

	fmt.Println("numPtr:", numPtr) //포인터 변수의 값(메모리 주소 출력)
	*numPtr = 24
	fmt.Println("numPtr:", *numPtr) //포인터 변수가 가리키는 메모리 주소에 들어있는 값 출력

	var num2 int = 58
	var num2Ptr *int = &num2
	fmt.Println("num2Ptr:", num2Ptr)  //포인터 변수의 값(메모리 주소 출력)
	fmt.Println("&num2:", &num2)      //일반 변수의 값 출력
	fmt.Println("num2Ptr:", *num2Ptr) //포인터 변수가 가리키는 메모리 주소에 들어있는 값 출력
}

func structPointerEx() {
	var rect1 Rectangle //구조체 인스턴스
	rect1.height = 10
	rect1.width = 20

	rect2 := new(Rectangle) //포인터 변수: var 키워드 없이 선언 + new로 메모리 할당

	fmt.Println("rect1:", rect1)
	fmt.Println("rect2:", rect2)

	rect3 := NewRectangle(20, 10)
	fmt.Println("rect3:", rect3)
	fmt.Println("rect3 area:", rectangleArea(rect3))
	fmt.Println("rect3 area(method):", rect3.area()) //구조체 메소드

	var s = Student
	s.p.greeting()
}

type Rectangle struct {
	width, height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height} //new는 제로값의 메모리 할당이고.. &는 값 넣은 메모리 할당
}

func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height //구조체 포인터의 값을읽을때도 .연산자 사용(구조체 인스턴스랑 똒같)
}

func (rect *Rectangle) area() int { //구조체 포인터에 대한 메서드
	return rect.width * rect.height
}

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello, my name is", p.name)
}

type Student struct {
	p      Person //(임베디드 필드) 구조체 안에 구조체 필드.(has-a 관계) (s.Person.name)
	Person        //이라면 (is-a 관계) (상속) Person 필드에 바로 접근 가능 (s.name)
	//만약!! Student도 Person 구조체와 같은 이름의 메소드가 있다면, s.Person.greeting은 Person의 greeting 호출, s.greeting은 Student의 greeting (오버라이딩) 호출
	school string
	grade  int
}
