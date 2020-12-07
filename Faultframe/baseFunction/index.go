package baseFunction

import (
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
)

func Index(fault *faultmodelstruct.Fault, common *commonStruct.Common) {



	LineFaultIndex(fault,common)
	TransFaultIndex(fault,common)
	//添加FaultIndex函数
}



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
		linefault.I = common.Ny + 1
		common.Ny += 2

	//观测变量映射
	}
}

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

