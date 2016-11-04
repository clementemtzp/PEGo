package main
import "fmt"
func main() {
	var a int;
	fmt.Scanf("%d", &a)
	for i :=a; i>0; i-- {
		for j:=0;j<i;j++{
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}