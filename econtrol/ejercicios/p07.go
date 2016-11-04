package main
import "fmt"
func main() {
	var n int
	var i int
	var j int
	fmt.Scanf("%d", &n)
	if n%2!=0{
		for i=1; i<=n/2+1; i++ {
			for j=1; j<=n-i; j++ {
				fmt.Print(" ")
			}
			for j=1; j<=2*i-1; j++ {
				fmt.Print("*")
			}
		fmt.Println(" ")
		}
		for i=(n-1)/2; i>=1; i-- {
			for j=1; j<=n-i; j++ {
				fmt.Print(" ");				
			}
			for j=1; j<=2*i-1; j++ {
				fmt.Printf("*");				
			}
		fmt.Println(" ")
		}		
	}
}