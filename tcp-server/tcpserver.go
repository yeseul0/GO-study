package main

import (
	"fmt"
	"net" //network 패키지(tcp, udp 통신)
)

func requestHandler(c net.Conn) { //net.Conn -> TCP 연결 객체
	data := make([]byte, 4096) //2^12 4096 바이트 슬라이스 생성(버퍼)

	for { //무한루프 -> 클라이언트가 보내는 데이터 계속 처리!!
		n, err := c.Read(data) //클라이언트에서 받은 데이터 읽기 n: 읽은 바이트 수
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) //데이터 출력

		_, err = c.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000") //8000 port에서 tcp 연결 대기 시작
	// "192.168.0.100:8000" 처럼 내 어떤 랜카드의 ip를 사용할지 콕 집어서 띄울 수도 있음!
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() //main끝 직전에 listener 닫음

	for { //계속 클라이언트 받음
		conn, err := ln.Accept() // 클라이언트 연결 오면 TCP 연결 객체 반환
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()   // 익명함수 끝나기 직전에 TCP 연결 닫음
			requestHandler(c) //패킷 처리함수를 고루틴으로 실행
		}(conn)
	}
}
