package main

import (
	"fmt"
	"net/rpc"
	"rpc-server/types"
)

// type Args struct {
// 	A, B int
// }

// type Reply struct {
// 	C int
// }

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	//동기 호출 (직접 기다림)
	args := &types.Args{1, 2}
	reply := new(types.Reply)
	err = client.Call("Calc.Sum", args, reply) //Calc.Sum 함수 호출. 값 리턴까지 기다림.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.C)

	//비동기 호출
	args.A = 4
	args.B = 5
	sumCall := client.Go("Calc.Sum", args, reply, nil) //Calc.Sum 함수를 고루틴으로 호출. *rpc.Call 객체가 들어옴.
	// 마지막 인자는 완료되면 신호 받을 채널인데, nil을 넣었으니 Go라이브러리가 알아서 채널 만들어서, rpc.Call.Done에 넣어줌
	<-sumCall.Done //채널에 값 들어올 때까지 대기
	fmt.Println(reply.C)
}
