package main 
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	userwel := "***************Welcome to the Terminal...***************"
	fmt.Println(userwel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ENter the rating for our Pizza: ")

	//comma-ok-
	 input, _ := reader.ReadString('\n')
	 fmt.Println("Thanks for rating = 5 ",input)
	 fmt.Printf("TYpe: %T \n",input)

}