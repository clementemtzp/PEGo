package main
import "fmt"
func main() {
	var num int;
	var suma int =1
	fmt.Println("Ingrese un numero")
	fmt.Scanf("%d", &num)
	for i := 2; i <= num; i++ {
		suma*=i
	}
	fmt.Println(num,"! =", suma)
}