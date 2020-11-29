package faultmodelstruct

type LineFault struct {
	BaseFault

	LineState   int32   //状态变量(索引)

	Load,Temperature,Health,I,U  int32   //观测变量(索引)

	//%(interfaceListCode)  int32   //接口变量

}