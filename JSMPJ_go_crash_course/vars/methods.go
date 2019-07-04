package main
import "fmt"
func main(){
  i:=10;
  passvalue(i);
  fmt.Println("Value of i after function call",i)
   fmt.Println("after pass value we are going to pass reference");
//    fmt.Println(==)
   passreference(&i);
   fmt.Println("Value of i after function ref fun call",i)
   sum(1, 2, 3, 4, 5)
   b:=del(6,3,"going") 
   fmt.Println("Function with return value result",b)
}
func passvalue(i int){
	fmt.Println("this is value of i",i);
	i=11
	fmt.Println("this is new value of i changed in func",i);
}
func passreference(i *int){
	*i=303
	fmt.Println("value of i after function refchange",*i);
}
func sum(values ...int){
	result:=0
	for _, v := range values {
		result+=v
	}
	fmt.Println("final result is", result);
}
func del(a int, b int, msg string) int , string{
	c:=a-b
	msg="coming"
	return c,msg
}