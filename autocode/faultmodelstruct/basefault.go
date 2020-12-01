package faultmodelstruct

// 设备结构体的基本结构体
type BaseFault struct {

	Name       string  // 故障设备名称
	DeviceType string  // 设备类型
	Pfault     float64 // T时刻故障概率
}

type Fault struct{
	LineFaults []LineFault //线路故障列表，列表中每个结构体用于对应线路的故障仿真
	TransFaults []TransFault


	LineFaultNames map[string]int32
	TransFaultNames map[string]int32
    //故障名与切片下标
}

