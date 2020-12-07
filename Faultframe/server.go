package main

import (
	"Faultframe/baseFunction"
	"Faultframe/commonStruct"
	"Faultframe/routines"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

// 二次系统

//SensorData 是接收数据的格式
type SensorData struct {
	Time           int
	ModelName      []string
	ObservationSum [][][]float64
}

//FaultP 是发送数据的格式
type FaultP struct {
	Time           int
	SamplingStatus [][][]bool
}

func main() {
	//启动server的监听器，并将连接丢给一个线程去做
	service := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 4)                            // set maxium request length to 128B to prevent flood attack
	var tempJSON []byte
	var sensor SensorData
	defer conn.Close() // close connection before exit
	for {
		//接收数据并转码，接收成功一组数据，处理完后再发送一组数据
		//循环监听，更新缓冲区
		readLen, err := conn.Read(request)
		checkError(err)
		//把缓冲数据读取出来放到tempJSON里面
		for j := 0; j < readLen; j++ {

			tempJSON = append(tempJSON, request[j])
		}
		//找到完整的数据段后将其print出来或者其他处理，接着向客户端发送数据，然后将处理过的数据段删除
		flag := -1
		for j := 0; j < len(tempJSON); j++ {
			if tempJSON[j] == 36 {
				//读取到截止符，执行断开连接的操作
				conn.Close()
			} else if tempJSON[j] == 35 {
				// 读取到分隔符
				// 执行数据处理
				err = json.Unmarshal(tempJSON[0:j], &sensor)
				checkError(err)
				fmt.Println(sensor)
				// 执行数据发送
				fp := "./grid/CEPRI36节点系统-改.xlsx"
				fault, common := baseFunction.Load(fp)
				baseFunction.Index(fault, common)
				baseFunction.Initial(fault, common)
				evidence := &commonStruct.Evidence{}
				environment := &commonStruct.Evidence_enviroment{}
				var time = 10.0
				tempStatus := routines.SimulateFault(fault, common, evidence, environment, time, sensor.ObservationSum)

				var SamplingStatus [][][]bool
				SamplingStatus = append(SamplingStatus, tempStatus)
				data := FaultP{
					sensor.Time,
					SamplingStatus,
				}
				b, err := json.Marshal(data)
				checkError(err)
				_, err = conn.Write(b)
				_, err = conn.Write([]byte("#"))
				checkError(err)
				flag = j
			}
		}
		tempJSON = tempJSON[flag+1:]
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
