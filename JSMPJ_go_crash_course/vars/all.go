package main
import "fmt"
type example struct{
	flag bool
	counter int16
	pi float32
}
func main(){
	aa:=2+5i
	fmt.Printf("%T\n",aa)
	var e1 example
	fmt.Printf("%+v\n",e1)
	e2:= example{
		flag:true,
		counter:10,
		pi:3.141592,
	}
	e3 := struct{
		flag bool
		counter int16
		pi float32
	}{
		flag:true,
		counter:20,
		pi:3.141516,
	}
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi",e2.pi)

	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi",e3.pi)
}