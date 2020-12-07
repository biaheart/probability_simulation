package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// SensorData 是发送数据格式
// 第一个是标记的时间点，第二个是观测量包含的模型例如是线路还是变压器，
// 和第三个观测量的顺序一一对应
type SensorData struct {
	Time           int
	ModelName      []string
	ObservationSum [][][]float64
}

// FaultP 是接收数据格式
// 第一个是标记的时间点，第二个是返回状态的采样数据
type FaultP struct {
	Time           int
	SamplingStatus [][][]bool
}

func main() {
	conn := buildChannel()
	//分次发送数据，#号作为每个数据的分隔符，$号作为结束通信的标志，由客户端发送完所有的包后发送
	go sendData(conn)
	receiveData(conn)

}

// buildChannel是建立与server的socket通信管道
func buildChannel() net.Conn {
	service := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	return conn
}

//sendData是模拟获得内存中的数据，用json转成字符编码发送
func sendData(conn net.Conn) {
	//分次发送数据，#号作为每个数据的分隔符，*号作为结束通信的标志，由客户端发送完所有的包后发送
	for i := 1; i < 10; i++ {
		var observationLineSum = [][][]float64{
			{
				{0.99},
				{1.00},
			},
			{
				{0.99},
				{0.99},
			},
		}
		data := SensorData{
			i,
			[]string{"line", "transfer"},
			observationLineSum,
		}
		b, err := json.Marshal(data)
		checkError(err)
		_, err = conn.Write(b)
		_, err = conn.Write([]byte("#"))
		checkError(err)
		if i == 9 {
			_, err = conn.Write([]byte("*"))
		}
	}
}

// receiveData是循环监听，更新缓冲区，
// 并操作数据（这一步可以丢给一个goroutine来完成，目前只是简单的print出来）
func receiveData(conn net.Conn) {
	//接收数据并转码
	// 缓冲区result
	result := make([]byte, 4)
	// json数据包临时变量
	var tempJSON []byte
	//json转码的格式变量
	var Probability FaultP
	//循环监听，更新缓冲区
	for {
		readLen, err := conn.Read(result)
		checkError(err)
		//把缓冲数据读取出来放到tempJSON里面
		for j := 0; j < readLen; j++ {

			tempJSON = append(tempJSON, result[j])
		}
		//找到完整的数据段后将其print出来或者其他处理，接着向客户端发送数据，然后将处理过的数据段删除
		flag := -1
		for j := 0; j < len(tempJSON); j++ {
			if tempJSON[j] == 35 {
				// 读取到分隔符
				// 执行数据处理
				err = json.Unmarshal(tempJSON[0:j], &Probability)
				checkError(err)
				fmt.Println(Probability)
				flag = j
			}
		}
		tempJSON = tempJSON[flag+1:]
	}
}

// checkError是检查异常的函数
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
