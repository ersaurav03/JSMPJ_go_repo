package main
import "fmt"
type vertex struct {
	X int 
	Y int
}
var (
	v1 = vertex{1,2}
	v2 = vertex{X:1}
	v3 = vertex{}
	p = &vertex{1,2} 
)
func main(){
   var a[4]string
   a[0]="JSMPJ"
   a[1]="Corporation"
   a[2]="ManakPur"
   fmt.Println(a)
   fmt.Println(a[0],a[1],a[2])
   fmt.Println(v1,v2,v3,p)

   b:=[7]int{1,2,3,4,5,6,7}
   fmt.Println(b)
   for i:=0; i<=len(b)-1; i++{
	   fmt.Println(b[i])
   }
   names := [4]string{
	   "John",
	   "Paul",
	   "george",
	   "ringo",
   }
   fmt.Println(names);
   a1 := names[0:2]
   a2 := names[1:3]
   a1[0]="YYY"
   fmt.Println(a1,a2);
   q := []int{1,2,3,4,5,6,7,8,9}
   fmt.Println(q)
   for i:=0;i<=len(q)-1;i++{
	   fmt.Println("variable present at position",i,"-->",q[i])
   }
   s := []struct {
	   i int
	   b bool
   }{
	   {2, true},
	   {3, false},
	   {4, true},
   }
   fmt.Println(s[0])
   s1:= []int{2,3,5,7,11,13}
   printSlice(s1)
   s1= s1[:4]
   printSlice(s1)
   s1= s1[2:]
   printSlice(s1)
   s1= s1[:0]
   printSlice(s1)
   var sz [5]int
   fmt.Println(sz, len(sz), cap(sz))
   ////Multidimesional array
   var array [3][3]int
   array[0] =[3]int{1,0,0}
   array[1] =[3]int{0,1,0}
   array[2] =[3]int{0,0,1}
   fmt.Println(array[0])
   fmt.Println(array[1])
   fmt.Println(array[2])
   const (
      a3 = iota + 1
      a4
      a5

   )
   const (
      Ldate = 1 << iota
      x2
      y2
      z2
      a7
   )
   fmt.Println(a3,a4,a5)
   fmt.Println(x2,y2,z2,a7)
   var fruit[5] int
   fruit[0]=1
   fruit[1]=2
   fruit[2]=3 
   fruit[3]=4
   fruit[4]=5
   fmt.Println(fruit)
   for i, fruits := range fruit {
      fmt.Println(i,fruits)
   }
}
func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s);

}