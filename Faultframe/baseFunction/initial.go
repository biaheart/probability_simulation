package baseFunction

import (
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
)

func Initial(fault *faultmodelstruct.Fault, common *commonStruct.Common) {




	//添加元件故障序列
	common.States = make([]bool, common.Nx)
	common.Evidences = make([]interface{}, common.Ny)
	common.PSensor = make([]float64, common.Ny)
	common.PTrans = make([]float64, common.Nx)
	//n := len(linefaults)





	//加入变量成功

}
