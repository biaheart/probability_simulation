package routines

import (
	"Faultframe/baseFunction"
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
	"fmt"
)

func SimulateFault(fault *faultmodelstruct.Fault,common *commonStruct.Common,evidence *commonStruct.Evidence,envoronment *commonStruct.Evidence_enviroment,time float64){

	for t:=float64(1) ;t<=time;{
		//TODO:插入socket通信的代码，获取一次系统的仿真数据，更新观测量
		baseFunction.SetLineEvidence(fault,common,evidence)
		baseFunction.SetTransEvidence(fault,common,evidence)//标幺值小于1.25数据才有效
		var observation_line_sum = [][][]float64{{{1.1, 1.11, 1.15, 1.11, 1.11,1.21,1.12,1.11,1.10,1.12},{1.08, 1.12, 1.09, 1.09, 1.12,1.12,1.15,1.16,1.17,1.19}},{{1.11, 1.11, 1.11, 1.11, 1.11,1.11, 1.11, 1.11, 1.11, 1.11},{1.09, 1.10, 1.11, 1.12, 1.15,1.12,1.11,1.11,1.12,1.13}}}
		var observation_trans_sum = [][][]float64{{{1.1, 1.11, 1.15, 1.11, 1.11},{1.08, 1.12, 1.09, 1.09, 1.12}},{{1.06, 1.06, 1.06, 1.06, 1.06},{1.06, 1.10, 1.10, 1.10, 1.10}}}
		observation := [][]float64{}
		observation_trans :=[][]float64{}
		var beliefP = 0.5                                            //先验信念
		var transP float64                                             //声明转移模型概率变量
		var GaussCoefficient = [][]float64{{1.1, 0.03}, {1.1, 0.03}} //高斯分布系数，行数和观测量行数要相同
		for i,_ :=range(observation_line_sum){
			observation = observation_line_sum[i]

			baseFunction.CalculateLineP(observation,beliefP,transP,GaussCoefficient,i,fault,common)
		}
		for t,_ :=range(observation_trans_sum){
			observation_trans = observation_trans_sum[t]

			baseFunction.CalculateTransP(observation_trans,beliefP,transP,GaussCoefficient,t,fault,common)
		}

		fmt.Println(t,"时刻故障状态：",common.States)
		common.PSensor_t = append(common.PSensor_t, common.PSensor)
		common.PSensor_t = append(common.PSensor_t,common.PTrans)
		t = t + 1
		//TODO：本时步故障仿真完成，插入socket通信的代码，通过socket向一次系统发送故障采样结果
	}
}