package faultmodelstruct

type TransFault struct {
	BaseFault

	TransState   int32

	Load,Health,I,U  int32

	//%(interfaceListCode)  int32
}

//className=LineFault name=linefault sheetName=Line  %slice=LineFaults

func TransFaultAdd(xlsx *excelize.File) []faultmodelstruct.TransFault{
	transfaultRows := xlsx.GetRows("Trans")
	if len(transfaultRows) == 0 {
		return []faultmodelstruct.TransFault{}
	}
	transfaultRows = transfaultRows[1:]
	TransFaults := make([]faultmodelstruct.TransFault,0)
	for i,_ := range transfaultRows{
		row := transfaultRows[i]
		if row[0][0] == '#' {
			continue
		}
		name:= row[0]
		devicetype := "TransFault"
		transfault:= faultmodelstruct.TransFault{}
		transfault.Name = name
		transfault.DeviceType = devicetype
		TransFaults= append(TransFaults,transfault)
	}
	return TransFaults
}


/////
//className=LineFault name=linefault sheetName=Line  %slice=LineFaults,typefaults=linefaults

func TransFaultIndex(fault *faultmodelstruct.Fault,common *commonStruct.Common){
    transfaults:= fault. TransFaults
    fault.TransFaultNames=map[string]int32{}
    for i,_ := range transfaults{
        transfault:= &transfaults[i]
        fault.TransFaultNames[transfault.Name]=int32(i)
    }
    //状态量、观测量索引建立
    for i,_ := range transfaults{
		transfault := &transfaults[i]

		transfault.TransState = common.Nx
		common.Nx += 1

		//状态变量索引映射
		transfault.Load = common.Ny
		transfault.Health = common.Ny + 1
		transfault.I = common.Ny + 2
		transfault.U = common.Ny + 3
		common.Ny += 4

		//观测变量映射
	}

}
/////
//className=LineFault name=linefault sheetName=Line  %slice=LineFaults,typefaults=linefaults

func SetTransEvidence(fault *faultmodelstruct.Fault,common *commonStruct.Common){
    //设置观测量的👈👈值
    transfaults := faults.TransFaults
	for i, _ := range(transfaults){

		common.Evidences[transfaults[i].Load] = evidence.Load
		common.Evidences[transfaults[i].Health] = evidence.Health
		common.Evidences[transfaults[i].I] = evidence.I
		common.Evidences[transfaults[i].U] = evidence.U

        //给观测变量序列赋值
	}
}



func TransTransModel(transP float64,index int,common *commonStruct.Common,faults *faultmodelstruct.Fault) float64 {

    var P_sensor float64
    if common.States[faults.TransFaults[index].TransState]==true{
        //插入TransP2
    }else{
        //插入TransP0

    }
    return P_sensor

    //计算完善的转移模型概率
}


func TransSensorModel(sensor []float64, GaussCoefficient [][]float64) [2]float64 {
	if len(sensor) != len(GaussCoefficient) {
		panic("观测量数量和概率分布系数数量不对应")
	}
	var sensorP = [2]float64{1, 0}
	for i, element := range sensor {
		sensorP[0] *= gauss(element, GaussCoefficient[i][0], GaussCoefficient[i][1])
	}
	sensorP[1] = 1 - sensorP[0] //这里粗暴地得到两个不同状态量条件下的获得对应观测量的概率，有待修改
	return sensorP
	//计算传感器模型概率
}



func CalculateTransP(observation [][]float64, beliefP float64, transP float64, GaussCoefficient [][]float64,index int,faults *faultmodelstruct.Fault,common *commonStruct.Common) ([]float64, []bool) {
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
	var samplingState []bool
	for i,_:=range(faultP) {
		q := rand.Float64()
		if q < faultP[i] {
			// common.States[faults.TransFaults[index].TransState] = true
			samplingState = append(samplingState, true)
		} else {
			samplingState = append(samplingState, false)
		}
	}
	//fmt.Println("faulttransP",faultP)
	return faultP, samplingState
	//返回某单个元件的时序故障概率
}

