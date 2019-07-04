package main
import ("fmt"
        "math")
var middle_name="kumar"
func main(){
	//Main Data types
	//String
	//bool
	//int int8 int16 int32 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128
        fmt.Println("JSMPJ Corporation")
	var name="M"
	var age int32=27
	var isCool=true
	fmt.Println(name,age,isCool)
	last_name:="Dhiman"
	fmt.Printf(last_name,"\n")
	fmt.Printf("%T",last_name);
	fmt.Println(middle_name)
	fmt.Println("find the type of strings %T--%T--%T",name,age,isCool)
	fmt.Println(math.Floor(2.7));
	fmt.Println(math.Ceil(3.4));
	fmt.Println(math.Sqrt(16));

}
