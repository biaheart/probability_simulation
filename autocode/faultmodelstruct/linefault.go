package faultmodelstruct

type LineFault struct {
	BaseFault

	LineState   int32   //״̬����(����)

	Load,Temperature,Health,I,U  int32   //�۲����(����)

	%(interfaceListCode)  int32   //�ӿڱ���

}
