//package main
//
//import (
//	"fmt"
//	"net"
//	"os"
//)
//
//func main() {
//	//通过tcp 协议链接 本机 8080端口
//	con, err := net.Dial("tcp", "127.0.0.1:8080")
//	//如果出现错误 说明链接失败
//	if err != nil {
//		fmt.Println("连接服务器端失败")
//		fmt.Println(err.Error())
//		os.Exit(0)
//	}
//	//记得关闭
//	defer con.Close()
//	var buffer []byte
//	go func() {
//		for{
//			n, _ := con.Read(buffer)
//			fmt.Println(string(buffer[:n]))
//		}
//	}()
//	for i:=float64(0);i<10;{
//		//开始向服务器端发送 hello
//		_, write_err := con.Write([]byte("Hello,I'm Fellow"))
//		//如果写入有问题 输出对应的错误信息
//		if write_err != nil {
//			fmt.Println(write_err.Error())
//		}
//		//如果没有问题。显示对应的写入长度
//		//fmt.Println(num)
//		i = i + 0.02
//	}
//	fmt.Println("通信完成")
//}
package main

import (
	"net"
	"fmt"
	"log"
	"os"
)
//发送信息
func sender(conn net.Conn) {
	words := "Hello I‘m Fellow-send!"
	conn.Write([]byte(words))
	fmt.Println("send over")

	//接收服务端反馈
	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)
	if err != nil {
		Log(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
		return
	}
	Log(conn.RemoteAddr().String(), "receive server back msg: ", string(buffer[:n]))

}
//日志
func Log(v ...interface{}) {
	log.Println(v...)
}

func main() {
	server := "127.0.0.1:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connection success")
	sender(conn)
}