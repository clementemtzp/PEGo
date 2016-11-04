package main
import "fmt"
func main() {
	var num int;
	var suma int =0
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		if i%3==0 || i%5==0{
			fmt.Println(i)
			suma+=i
		}
	}
	fmt.Println("Suma es:", suma)
}