package main
// C:\Users\HP\Desktop\golang\2-variables\main.go
import "fmt"

func main(){
	var username string = "gnanesh" 
	fmt.Println("Username : ",username)
	fmt.Printf("variables Type: %T \n",username)

	var isloggedIn bool = true
	fmt.Println("Login T/F : ",isloggedIn)
	fmt.Printf("variable Type: %T \n",isloggedIn)

	var snum uint8 = 255
	fmt.Println(snum)
	fmt.Printf("Variable Type: %T \n",snum)

	var sfloat float64 = 27904.498757957759757959
	fmt.Println(snum)
	fmt.Printf("Variable Type: %T \n",sfloat)


	// defualt values and (alias) 
	var notinalvar1 int 
	fmt.Println(notinalvar1)
	fmt.Printf("Variable Type: %T \n",notinalvar1)

	//implicit 
	var web = "Gnanesh-16"
	fmt.Println(web)

	//no-var-style
	name  := "Gnanesh "
	num := 3000000
	fmt.Println(name,num)


}