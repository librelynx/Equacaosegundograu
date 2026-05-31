/******************************************************************************
O problema da equação do segundo grau
*******************************************************************************/
package main
import ("fmt"
	"math")

func main() {
	//declara as variáveis
	var a float64 = 1
	var b float64 = 2
	var c float64 = 1

	//calcula as raízes
	delta := b*b - 4*a*c
	x1 := (-b + math.Sqrt(float64(delta))) / (2 * a)
	x2 := (-b - math.Sqrt(float64(delta))) / (2 * a)

	//imprime os resultados
	fmt.Println("x1 =", x1)
	fmt.Println("x2 =", x2)
}
