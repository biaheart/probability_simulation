package faultmodelstruct

type TransFault struct {
	BaseFault

	TransState   int32   //״̬����(����)

	Load,Health,I,U  int32   //�۲����(����)

	%(interfaceListCode)  int32   //�ӿڱ���

}
