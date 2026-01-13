package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8000") //TCP 프로토콜, 127.0.0.1:8000서버에 연결!!
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	go func(c net.Conn) {
		data := make([]byte, 4096) //4096바이트 버퍼
		for {
			n, err := c.Read(data) //서버에서 받은 데이터. n에는 바이트 수
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data[:n]))
			time.Sleep(1 * time.Second)
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			s := "hello" + strconv.Itoa(i) //hello0,hello1, ...
			_, err := c.Write([]byte(s))   //서버로 데이터 보내기
			if err != nil {
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	fmt.Scanln()
}
