package baseFunction

import (
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"

)
func LineTransModel(transP float64,index int,common *commonStruct.Common,faults *faultmodelstruct.Fault) float64 {

    var P_sensor float64
    if common.States[faults.LineFaults[index].LineState]==true{
		P_sensor =0.2        //插入LineP2
    }else{
		P_sensor=0.85        //插入LineP0

    }
    return P_sensor

	//计算完善的转移模型概率
}

func TransTransModel(transP float64,index int,common *commonStruct.Common,faults *faultmodelstruct.Fault) float64 {

    var P_sensor float64
    if common.States[faults.TransFaults[index].TransState]==true{
		P_sensor =0.05        //插入TransP2
    }else{
		P_sensor=0.9        //插入TransP0

    }
    return P_sensor

	//计算完善的转移模型概率
}

