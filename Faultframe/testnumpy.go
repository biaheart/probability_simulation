package main

import "fmt"

func main(){
	test:=[]float64 {1,3,5,7}
	for i,_:= range (test){
		if 3<test[i] {
			fmt.Println(test[i])
		}

	}

}