package main

import "fmt"

func main() {
    var arr [32]int32
    fmt.Println(arr[0])
    fmt.Println(arr[1:3])

    // var arr2 = [1]int{123}
    // fmt.Println(arr2[0])
    // fmt.Println(arr2[:]) 
	
	var arr2 = [3]int{0, 1, 2} 
    fmt.Println(arr2[0:3]) 
	fmt.Println(&arr2[0])
    fmt.Println(&arr2[1])
    fmt.Println(&arr2[2])
}
