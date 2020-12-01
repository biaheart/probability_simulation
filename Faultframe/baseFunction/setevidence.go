package baseFunction

import (
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
)



func SetLineEvidence(faults *faultmodelstruct.Fault,common *commonStruct.Common,evidence *commonStruct.Evidence){
    //设置观测量的👈👈值
    linefaults := faults.LineFaults
	for i, _ := range(linefaults){

		common.Evidences[linefaults[i].Load] = evidence.Load
		common.Evidences[linefaults[i].I] = evidence.I

	//给观测变量序列赋值
	}
}

func SetTransEvidence(faults *faultmodelstruct.Fault,common *commonStruct.Common,evidence *commonStruct.Evidence){
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

