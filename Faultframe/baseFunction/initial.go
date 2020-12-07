package baseFunction

import (
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
)

func Initial(fault *faultmodelstruct.Fault, common *commonStruct.Common) {




	linefaults:= fault.LineFaults
	transfaults:= fault.TransFaults
	//添加元件故障序列
	common.States = make([]bool, common.Nx)
	common.Evidences = make([]interface{}, common.Ny)
	common.PSensor = make([]float64, common.Ny)
	common.PTrans = make([]float64, common.Nx)
	//n := len(linefaults)





	for i,_ := range(linefaults){
		linefault:=linefaults[i]
		common.Evidences[linefault.Load] = 0
		common.Evidences[linefault.I] = 0
		common.States[linefault.LineState] = false
	}
	for i,_ := range(transfaults){
		transfault:=transfaults[i]
		common.Evidences[transfault.Load] = 0
		common.Evidences[transfault.Health] = 0
		common.Evidences[transfault.I] = 0
		common.Evidences[transfault.U] = 0
		common.States[transfault.TransState] = false
	}
	//加入变量成功

}
