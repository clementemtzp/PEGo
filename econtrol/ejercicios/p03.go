package main
import "fmt"
func main() {
	var a int;
	var b int;
	var suma int =1
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	for i :=0; i <b; i++ {
		suma*=a
	}
	fmt.Println(a,"^",b,"=",suma)
}