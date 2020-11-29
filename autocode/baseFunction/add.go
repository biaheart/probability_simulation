package baseFunction

import (
	"Faultframe/commonStruct"
	"Faultframe/faultmodelstruct"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)
func Load(fp string)(*faultmodelstruct.Fault,*commonStruct.Common){
	xlsx, err := excelize.OpenFile(fp)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fault := faultmodelstruct.Fault{}

	fault.LineFaults= LineFaultAdd(xlsx)
	fault.TransFaults= TransFaultAdd(xlsx)
	fault.LineFaults= LineFaultAdd(xlsx)
	fault.TransFaults= TransFaultAdd(xlsx)
    //调用FaultAdd函数

	common := commonStruct.Common{}
	return &fault,&common
}

