package main
import "fmt"
func main() {
	var n int
	var i int
	var j int
	var x int=1
	fmt.Scanf("%d", &n)
	for i=0; i<n; i++{
		for j=0; j<=i; j++{
			fmt.Print(x," ")
			x+=1
		}
		fmt.Println("")
	}
}