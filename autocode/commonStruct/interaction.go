package commonStruct

type Evidence struct{
	I []float64 `json:"I"`
	U []float64	`json:"U"`
	Load[]float64 `json:"Load"`
}

type Evidence_protect struct {

}

type Evidence_enviroment struct {
	Health []float64	`json:"Health"`
	Temparature []float64 `json:"Temparature"`
}