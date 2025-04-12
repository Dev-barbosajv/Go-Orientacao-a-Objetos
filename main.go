package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float32
} 



func main() {
	ContaDoJoao := ContaCorrente{titular: "Joao Vitor", numeroAgencia: 1234, numeroConta: 3322, saldo: 124.50}
	ContaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(ContaDoJoao)
	fmt.Println(ContaDaBruna) 
}
	