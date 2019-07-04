package main
import "fmt"
func main(){
	x:=50
	y:=10
	if (x < y) {
		fmt.Printf("%d is less than %d\n",x,y)
	} else
	{
		fmt.Println("Bye")
	}
	emails := make(map[string]string)
	emails2 :=map[string]string{"bob":"bob@gmail.com","sharon":"sharon@gmail.com"}
	emails["bob"]="bob@gmail.com"
	emails["sharon"]="sharon@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))

	fmt.Println(emails2)
	fmt.Println(len(emails2))
	delete(emails,"bob")
	
	fmt.Println(emails)
	fmt.Println(len(emails))
	color:="red"
	switch color{
	case "green":
		fmt.Println("color is green")
	case "reda":
		fmt.Println("color is red")
	default:
		fmt.Println("color is default")
	}
	ids := []int{11,22,33,44,55}
	for i, id := range ids {
		fmt.Printf("%d - ID : %d\n", i, id)
	}
	for _, id := range ids {
		fmt.Printf(" ID : %d\n", id)
	}
    emailsss := map[string]string{"saurav":"Saurav@gmailcom","kumar":"kumar@gmail.com"}
    for k, v := range emailsss{
		fmt.Printf("name:-- %s ,emails address:-%s\n",k,v)
	}

	for names := range emailsss{
		fmt.Printf("name:-- %s\n ",names)
	}

	fmt.Println("----JSMPJ pointersss starts---")
	ay:= 5
	by:= &ay

	fmt.Println(ay,*by)
	*by=20
	fmt.Println(ay,*by)
	
	 

}