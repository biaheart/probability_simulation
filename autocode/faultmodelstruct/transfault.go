package faultmodelstruct

type TransFault struct {
	BaseFault

	TransState   int32   //状态变量(索引)

	Load,Health,I,U  int32   //观测变量(索引)

	%(interfaceListCode)  int32   //接口变量

}
