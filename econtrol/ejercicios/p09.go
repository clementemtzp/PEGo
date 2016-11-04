package main
import "fmt"
func main() {
	var n int
	var i int
	var j int
	var c int=0
	var t int
	fmt.Scanf("%d", &n)
	for i=0; i<=n; i++ {
		t=c;
		c+=i;
		for j=c; j>t; j-- {
			fmt.Print(j," ")
		}
		fmt.Println("")
	}
}