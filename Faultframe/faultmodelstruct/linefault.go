package faultmodelstruct

type LineFault struct {
	BaseFault

	LineState   int32

	Load,I  int32

	//%(interfaceListCode)  int32
}
