package main

import (
	"net/http"
)

func ex1() {
	//라우팅. / 주소로 오면, 이 함수를 실행해!
	http.HandleFunc("/",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte("hello, go world!")) //웹 브라우저 응답
		}) //경로 접속 시 실행 함수 설정

	http.ListenAndServe(":80", nil) // 80번 포트에서 웹서버 실행. nil -> 기본 mux 사용(DefaultServeMux)
}

// ListenAndServe에 멀티플렉서 핸들러를 넣음
func ex2() {
	mux := http.NewServeMux() //HTTP 요청 멀티플렉서 생성
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, gogo!"))
	})

	mux2 := http.NewServeMux() //HTTP 요청 멀티플렉서 생성
	mux2.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, hoho!"))
	})

	go http.ListenAndServe(":8000", mux) //고루틴 실행
	http.ListenAndServe(":80", mux2)     //메인 실행
	//ListenAndServe 줄에서 블로킹 되거든 ^^
}
