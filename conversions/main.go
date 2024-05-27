package main

//conversion string use (strconv),
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	fmt.Println("****** Welcome to pizza app******")
	fmt.Println(" Rate on Scale of 1-5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating , ",input)

	numrate, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to Rating: ", numrate + 1)

	}

}