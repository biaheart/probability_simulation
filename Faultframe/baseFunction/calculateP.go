package baseFunction

import (
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
	"fmt"
	"math/rand"
)

func CalculateLineP(observation [][]float64, beliefP float64, transP float64, GaussCoefficient [][]float64,index int,faults *faultmodelstruct.Fault,common *commonStruct.Common) []float64 {
	var faultP []float64        //存储时序故障概率
	var posterior, priorP float64 //后验概率,先验概率
	var t, i int
	transP = LineTransModel(transP,index,common,faults) //根据转移模型获得转移概率
	posterior = beliefP         //将先验信念作为零时刻的后验概率
	for t = 0; t < len(observation[0]); t++ {
		var sensor []float64   //获取某一时刻的所有观测量的值
		var sensorP [2]float64 //存储传感器模型概率
		for i = 0; i < len(observation); i++ {
			sensor = append(sensor, observation[i][t])
			//fmt.Println("sensor",sensor)
		}
		priorP = transP*posterior + (1-transP)*(1-posterior)                          //更新t时刻的先验概率
		sensorP = LineSensorModel(sensor, GaussCoefficient)                               //根据传感器模型获得传感器概率
		posterior = sensorP[0] * priorP / (sensorP[0]*priorP + sensorP[1]*(1-priorP)) //更新t时刻的后验概率
		faultP = append(faultP, posterior)
		//fmt.Println("faultP",faultP)
	}
	for i,_:=range(faultP) {
		q := rand.Float64()
		if q < faultP[i]{
			common.States[faults.LineFaults[index].LineState] = true
		}
	}
	fmt.Println("faulttransP",faultP)
	return faultP
	//返回某单个元件的时序故障概率
}

func CalculateTransP(observation [][]float64, beliefP float64, transP float64, GaussCoefficient [][]float64,index int,faults *faultmodelstruct.Fault,common *commonStruct.Common) []float64 {
	var faultP []float64        //存储时序故障概率
	var posterior, priorP float64 //后验概率,先验概率
	var t, i int
	transP = TransTransModel(transP,index,common,faults) //根据转移模型获得转移概率
	posterior = beliefP         //将先验信念作为零时刻的后验概率
	for t = 0; t < len(observation[0]); t++ {
		var sensor []float64   //获取某一时刻的所有观测量的值
		var sensorP [2]float64 //存储传感器模型概率
		for i = 0; i < len(observation); i++ {
			sensor = append(sensor, observation[i][t])
			//fmt.Println("sensor",sensor)
		}
		priorP = transP*posterior + (1-transP)*(1-posterior)                          //更新t时刻的先验概率
		sensorP = TransSensorModel(sensor, GaussCoefficient)                               //根据传感器模型获得传感器概率
		posterior = sensorP[0] * priorP / (sensorP[0]*priorP + sensorP[1]*(1-priorP)) //更新t时刻的后验概率
		faultP = append(faultP, posterior)
		//fmt.Println("faultP",faultP)
	}
	for i,_:=range(faultP) {
		q := rand.Float64()
		if q < faultP[i]{
			common.States[faults.TransFaults[index].TransState] = true
		}
	}
	//fmt.Println("faulttransP",faultP)
	return faultP
	//返回某单个元件的时序故障概率
}

