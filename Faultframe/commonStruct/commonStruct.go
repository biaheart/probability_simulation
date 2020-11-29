package commonStruct

type Common struct {
	States    []bool        //所有故障模型的状态变量列表
	Evidences []interface{} //所有观测变量的列表
	PSensor   []float64     //传感模型求得的概率
	PTrans    []float64     //转移模型求得的概率
	//Ptransfer []float64 //所有转移模型的概率值
	//Psensor []float64 //所有传感器模型的概率值
	PSensor_t [][]float64 //历史概率记录
	PTrans_t  [][]float64 //历史概率记录
	Nx,Ny     int32       // 状态量以及观测量个数
	faultP [] float64    //故障概率
}

