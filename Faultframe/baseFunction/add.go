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
	//调用FaultAdd函数
	common := commonStruct.Common{}
	return &fault,&common
}
func LineFaultAdd(xlsx *excelize.File) []faultmodelstruct.LineFault{
	linefaultRows := xlsx.GetRows("Line")
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

func TransFaultAdd(xlsx *excelize.File) []faultmodelstruct.TransFault{
	transfaultRows:= xlsx.GetRows("Trans")
	if len(transfaultRows) == 0 {
		return []faultmodelstruct.TransFault{}
	}
	transfaultRows = transfaultRows[1:]
	TransFaults := make([]faultmodelstruct.TransFault,0)
	for i,_ := range transfaultRows{
		row := transfaultRows[i]
		if row[0][0] == '#' {
			continue
		}
		name:= row[0]
		devicetype := "TransFault"
		transfault:= faultmodelstruct.TransFault{}
		transfault.Name = name
		transfault.DeviceType = devicetype
		TransFaults= append(TransFaults,transfault)
	}
	return TransFaults
}

