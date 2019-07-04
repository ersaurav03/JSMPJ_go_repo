package main
import "fmt"
func main(){
	var a int
	var b int
	fmt.Println("enter the value for A")
	fmt.Scanln(&a)
	fmt.Println("enter the value for B")
	fmt.Scanln(&b)
	fmt.Println("Entered the value for A",a)
	fmt.Println("Entered the value for B",b)
    fmt.Println("Addition of and b",add(a,b))
}
func add(a, b int) int{
	c:=a+b
	return c
}