package main
import (
	"fmt"
	"strconv"
)
type names struct{
	name string
	age int
}
func (n names) greet() string{
	return "hello name is "+n.name+ "age is"+strconv.Itoa(n.age)
}
func main(){
name1:=names{"Saurav",27}
 fmt.Println(name1.greet())
}