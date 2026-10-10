// declarar pacote principal
package main

//importar a função fmt
import "fmt"

// Valor de ebulição da água em Kelvin
const ebulicaoK float64 = 373.15

// Função principal
func main() {

	var ebulicaoC = ebulicaoK - 273.15

	fmt.Println("A temperatura de ebulição da água em K = ", ebulicaoK)
	fmt.Println("A temperatura de ebulição da água em °C = ", ebulicaoC)

}
