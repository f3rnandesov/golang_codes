package main

import "fmt"

func main(){
	var N int
	fmt.Print("Digite um número inteiro positivo: ")
	_, err := fmt.Scan(&N)

	if err != nil || N <= 0{
		fmt.Println("Por favor, digite número inteiro positivo.")
		return 
	}

	conta := N * (N + 1) / 2
	fmt.Printf("Resultado: %d\n", conta)
	fmt.Println("Fórmula usada: N x (N + 1) / 2")


}