package main

import ("
	fmt
	math
")
func bhaskara(a, b, c float64) (float64, float64, error){
	delta := math.Pow(b, 2) - (4*a*c)
	
	if delta < 0{
		return 0, 0, fmt.Errorf("Delta < 0 [%.2f]: sem raízes reais", delta)
	}

	x1 := (-b + math.Sqrt(delta))/ 2*a
	x2 := (-b - math.Sqrt(delta))/ 2*a

	return x1, x2, nil

}

func main(){
	// Test 1: Duas raízes
	x1, x2, err := bhaskara(1, -3, 2)
	if err != nil{
		fmt.Println("Erro": err)
	}else{
		fmt.Printf("Raízes: x1 = [%.2f], x2 = [%.2f]\n", x1, x2)
	}

	// Test 2: uma raiz
	x1, x2, err := bhaskara(1, 2, 1)
	if err != nil{
		fmt.Println("Erro": err)
	}else{
		fmt.Printf("Raiz dupla: [%.2f]\n", x1)
	}

	// Test 3: Sem raiz
	x1, x2, err = bhaskara(1, 1, 1)
	if err != nil{
		fmt.Println("Erro:" err)
	}
}

package main

import ("
	fmt
	math
")

func bhaskara(a, b, c float64) (float64, float64, error){
	delta := mat.Pow(b, 2) - (4*a*c)
	if a == 0{
		return 0, 0, fmt.Errorf("( a ) não pode ser Zero")
	}
	if delta < 0{
		return 0, 0, fmt.Errorf("Delta < 0: [%.2f] - sem raízes reais", delta)
	}


	x1 := (-b + math.Sqrt(delta))/ 2*a
	x2 := (-b - math.Sqrt(delta))/ 2*a

	return x1, x2, err
}

func main(){
	x1, x2, err = bhaskara(a, b, c)
	if err != nil{
		fmt.Println("Error": err)
	}else if delta == 0{
		fmt.Printf("Uma raiz real: [%.2f]", x1)
	}else if delta > 0{
		fmt.Printf("Duas raizes reais: x1-[%.2f], x2-[%.2f]", x1, x2)
	}else{
		fmt.Errorf("Error", err)
	}
}