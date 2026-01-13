package main //다른 컴터에서 내 함수를 실행하게 해주는 서버

import (
	"fmt"
	"net"
	"net/rpc"
	"rpc-server/types"
)

type Calc int //RPC 서버에 등록하기 위해 임의의 타입으로 정의?

type Args struct { //매개변수
	A, B int
}

type Reply struct { // 리턴값
	C int
}

// Calc 타입에 소속된 메소드 Sum. *Calc 통해서만 호출할 수 있음
func (c *Calc) Sum(args types.Args, reply *types.Reply) error { //원격 실행 시켜줄 함수 signature
	fmt.Printf("요청 (A: %d, B: %d)\n", args.A, args.B)
	reply.C = args.A + args.B //매개변수 두 값을 더하여, 리턴 구조체에 넣어줌(포인터 접근임!)
	fmt.Printf("리턴 %d\n", reply.C)
	return nil
}

func main() {
	rpc.Register(new(Calc))               //Calc 타입의 인스턴스를 RPC 서버에 등록!-> client는 Calc.sum으로 호출할 수 있게됨
	ln, err := net.Listen("tcp", ":6000") //TCP 프로토콜, 6000번 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() //main 함수가 종료되기 직전에 연결대기 닫음

	for {
		conn, err := ln.Accept() //클라이언트가 연결되면 TCP 연결 객체 (net.Conn) 리턴.
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			rpc.ServeConn(c) //c로 들어오는 바이트 stream 읽어서, Calc.Sum 호출하는구나! 파악후 실행하고 결과값 바이트로 쏴줌
		}(conn)
		// 왜 *net.Conn을 주지 않고도 괜찮지?
		// net.Conn은 실제 타입 정보(*net.TCPConn)&실제 소켓 객체 주소를 담고있는 interface임!!! 애초에 내부가 포인터 2개 ㅋ
	}
}
