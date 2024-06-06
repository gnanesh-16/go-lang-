package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Time from golang__..")
	presentTime := time.Now()
	fmt.Fprint(os.Stdout, []any{presentTime}...)
	fmt.Println(presentTime.Format("01-02-2006")) //need to give this same formate to get a real date standered syntax no exception
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	// fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	crtdate := time.Date(2020, time.May, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(crtdate)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

}
