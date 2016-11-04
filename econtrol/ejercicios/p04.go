package main
import "fmt"
func main() {
	var a int;
	var b int;
	var suma int =0
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	for i :=0; i <b; i++ {
		suma+=a
	}
	fmt.Println("El producto es",suma)
}