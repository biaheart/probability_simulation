package routines

import (
	"Faultframe/baseFunction"
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
	"fmt"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//SimulateFault 是故障仿真程序
func SimulateFault(fault *faultmodelstruct.Fault, common *commonStruct.Common, evidence *commonStruct.Evidence, envoronment *commonStruct.Evidence_enviroment, time float64, observationLineSum [][][]float64) [][]bool {

	//TODO:插入socket通信的代码，获取一次系统的仿真数据，更新观测量
	baseFunction.SetLineEvidence(fault, common, evidence)
	baseFunction.SetTransEvidence(fault, common, evidence) //标幺值小于1.25数据才有效
	var observationTransSum = [][][]float64{
		{
			{1.10, 1.11, 1.15, 1.11, 1.11, 1.02, 1.03, 1.01, 1.04, 1.02, 1.01, 0.99, 1.05},
			{1.08, 1.12, 1.09, 1.09, 1.12, 1.08, 1.06, 1.09, 1.05, 1.06, 1.07, 1.05, 1.03},
		},
		{
			{1.06, 1.06, 1.06, 1.06, 1.06, 1.03, 1.02, 1.02, 1.02, 1.02, 1.01, 0.99, 0.98, 1.00, 1.02, 1.03, 1.01, 1.08, 1.06, 1.09, 1.05, 1.06, 1.03, 1.01, 1.04, 1.02, 1.01, 0.99, 1.05, 1.05, 1.06, 1.04, 1.07, 1.03, 1.03, 1.02, 1.02, 1.02, 1.03, 1.04, 1.06, 1.08, 1.10, 1.09, 1.09, 1.07, 1.05, 1.08, 1.10, 1.09, 1.09},
			{1.06, 1.10, 1.10, 1.10, 1.10, 1.01, 0.99, 0.98, 0.99, 1.00, 1.01, 1.00, 1.02, 1.03, 1.05, 1.05, 1.06, 1.04, 1.07, 1.03, 1.03, 1.02, 1.02, 1.02, 1.02, 0.99, 0.98, 0.97, 0.99, 1.03, 1.05, 1.03, 0.99, 0.98, 0.97, 0.99, 1.01, 1.03, 1.04, 1.06, 1.08, 1.10, 1.09, 1.11, 1.11, 1.11, 1.11, 1.11, 1.14, 1.12, 1.13},
		},
	}
	observation := [][]float64{}
	observationTrans := [][]float64{}
	var beliefP = 0.5  //先验信念
	var transP float64 //声明转移模型概率变量
	var SamplingStatus [][]bool
	var GaussCoefficient = [][]float64{{1.1, 0.03}, {1.1, 0.03}} //高斯分布系数，行数和观测量行数要相同
	for i := range observationLineSum {
		observation = observationLineSum[i]

		var y, samplingState = baseFunction.CalculateLineP(observation, beliefP, transP, GaussCoefficient, i, fault, common)
		var address = "line" + strconv.Itoa(i) + ".jpg"
		var x []float64
		for j := 1; j <= len(observation[0]); j++ {
			x = append(x, float64(j))
		}
		Plot(x, y, address)
		// fmt.Println("line" + strconv.Itoa(i) + "state")
		// fmt.Println(samplingState)
		SamplingStatus = append(SamplingStatus, samplingState)
	}
	for t := range observationTransSum {
		observationTrans = observationTransSum[t]

		var y, samplingState = baseFunction.CalculateTransP(observationTrans, beliefP, transP, GaussCoefficient, t, fault, common)
		var address = "trans" + strconv.Itoa(t) + ".jpg"
		var x []float64
		for j := 1; j <= len(observationTrans[0]); j++ {
			x = append(x, float64(j))
		}
		Plot(x, y, address)
		fmt.Println("trans" + strconv.Itoa(t) + "state")
		fmt.Println(samplingState)
	}

	common.PSensor_t = append(common.PSensor_t, common.PSensor)
	common.PSensor_t = append(common.PSensor_t, common.PTrans)
	//TODO：本时步故障仿真完成，插入socket通信的代码，通过socket向一次系统发送故障采样结果
	return SamplingStatus
}

//Plot 是画折线图的函数
func Plot(x, y []float64, address string) {
	p, _ := plot.New()
	points := plotter.XYs{}
	for i := 0; i < len(x); i++ {
		points = append(points, plotter.XY{X: x[i], Y: y[i]})
	}
	plotutil.AddLines(p, points)
	p.Save(8*vg.Inch, 4*vg.Inch, address)
}
