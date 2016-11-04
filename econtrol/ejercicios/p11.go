package main
import "fmt"
func main() {
	var x int=0
	var y int=1
	var z int
	var n int
	fmt.Println("Dame un numero")
	fmt.Scanf("%d", &n)
		fmt.Println("0")
		fmt.Println("1")
	for i:=1;i<=n; i++ {
		z=x+y
		fmt.Println(z)
		x=y
		y=z
	}
}