package main

import (
	"fmt"
	"net"
)

func main() {

	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer conn.Close()

	//发送数据
	conn.Write([]byte("are you ok?"))
}
