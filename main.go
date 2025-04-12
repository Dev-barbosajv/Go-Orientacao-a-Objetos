package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float32
} 



func main() {
	var titular string = "Joao Vitor"
	var numeroAgencia int = 589
	var numeroConta int = 12345
	var saldo float32 = 125.50

	fmt.Println(titular, numeroAgencia, numeroConta, saldo) 
}
	