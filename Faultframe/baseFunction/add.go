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
	fault.LineFaults = LineFaultAdd(xlsx)
	fault.TransFaults= TransFaultAdd(xlsx)
	common := commonStruct.Common{}
	return &fault,&common
}




