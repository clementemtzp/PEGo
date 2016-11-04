package main
import "fmt"
func main() {
	for i := 0; i <=6; i++ {
		switch i {
		case 0: fmt.Println("Cero")
		case 1: fmt.Println("Uno")
		case 2: fmt.Println("Dos")
		case 3: fmt.Println("Tres")
		case 4: fmt.Println("Cuatro")
		case 5: fmt.Println("Cinco")
		default: fmt.Println("Otro numero no declarado")
		}
	}
}