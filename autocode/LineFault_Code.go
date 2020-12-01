package faultmodelstruct

type LineFault struct {
	BaseFault

	LineState   int32   //状态变量(索引)

	Load,Temperature,Health,I,U  int32   //观测变量(索引)

	%(interfaceListCode)  int32   //接口变量

}

//className=LineFault name=linefault sheetName=Line  %slice=LineFaults

func LineFaultAdd(xlsx *excelize.File) []faultmodelstruct.LineFault{
	linefaultRows,_ := xlsx.GetRows("Line")
	if len(linefaultRows) == 0 {
		return []faultmodelstruct.LineFault{}
	}
	linefaultRows = linefaultRows[1:]
	LineFaults := make([]faultmodelstruct.LineFault,0)
	for i,_ := range linefaultRows{
		row := linefaultRows[i]
		if row[0][0] == '#' {
			continue
		}
		name:= row[0]
		devicetype := "LineFault"
		linefault:= faultmodelstruct.LineFault{}
		linefault.Name = name
		linefault.DeviceType = devicetype
		LineFaults= append(LineFaults,linefault)
	}
	return LineFaults
}


/////
//className=LineFault name=linefault sheetName=Line  %slice=LineFaults,typefaults=linefaults

func LineFaultIndex(fault *faultmodelstruct.Fault,common *commonStruct.Common){
    linefaults:= fault. LineFaults
    fault.LineFaultNames=map[string]int32{}
    for i,_ := range linefaults{
        linefault:= &linefaults[i]
        fault.LineFaultNames[linefault.Name]=int32(i)
    }
    //状态量、观测量索引建立
    for i,_ := range linefaults{
		linefault := &linefaults[i]

		linefault.LineState = common.Nx
		common.Nx += 1

		//状态变量索引映射
		linefault.Load = common.Ny
		linefault.Temperature = common.Ny + 1
		linefault.Health = common.Ny + 2
		linefault.I = common.Ny + 3
		linefault.U = common.Ny + 4
		common.Ny += 5

		//观测变量映射
	}

}
/////
//className=LineFault name=linefault sheetName=Line  %slice=LineFaults,typefaults=linefaults

func SetLineEvidence(fault *faultmodelstruct.Fault,common *commonStruct.Common){
    //设置观测量的👈👈值
    linefaults := faults.LineFaults
	for i, _ := range(linefaults){

		common.Evidences[linefaults[i].Load] = evidence.Load
		common.Evidences[linefaults[i].Temperature] = evidence.Temperature
		common.Evidences[linefaults[i].Health] = evidence.Health
		common.Evidences[linefaults[i].I] = evidence.I
		common.Evidences[linefaults[i].U] = evidence.U

        //给观测变量序列赋值
	}
}



func LineTransModel(transP float64,index int,common *commonStruct.Common,faults *faultmodelstruct.Fault) float64 {

    var P_sensor float64
    if common.States[faults.LineFaults[index].LineState]==true{
        //插入LineP2
    }else{
        //插入LineP0

    }
    return P_sensor

    //计算完善的转移模型概率
}


func LineSensorModel(sensor []float64, GaussCoefficient [][]float64) [2]float64 {
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
	//fmt.Println("faulttransP",faultP)
	return faultP
	//返回某单个元件的时序故障概率
}

