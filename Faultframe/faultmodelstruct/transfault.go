package faultmodelstruct

type TransFault struct {
	BaseFault

	TransState   int32

	Load,Health,I,U  int32

	//%(interfaceListCode)  int32
}
