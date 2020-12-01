package routines

import (
	"Faultframe/baseFunction"
	"Faultframe/commonStruct"
	"testing"
)

func TestSimulateFault(t *testing.T) {
	fp := "../grid/CEPRI36节点系统-改.xlsx"
	fault, common := baseFunction.Load(fp)
	baseFunction.Index(fault, common)
	baseFunction.Initial(fault, common)
	evidence := &commonStruct.Evidence{}
	environment := &commonStruct.Evidence_enviroment{}
	var time = 10.0
	SimulateFault(fault, common, evidence, environment, time)
}
