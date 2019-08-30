package main

import "fmt"

func zeroVal(ival string) {
	ival = "valor função normal"
	// ival = 0
}

func zeroPtr(iptr *string) {
	*iptr = "valor funcao ponteiro"
	// *iptr = 0
}

func main() {

	// i := 0
	i := "esse é o meu valor chamando a função"
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	fmt.Println("pointer:", &i)
}
