package main /*Compiler should know were actually program should start thats the reseaon main needt o be declared  */
import "fmt"

func main(){
	fmt.Println("Go -lang upskil.. ll... ")
	// basics()
	array()
	array_2()
	array_3()
	array_4()
}


func array(){ //intiall 0 
	var arr1 [3]int32
	fmt.Println(arr1)

	arr1[1] = 123
	fmt.Println(arr1[0])
	fmt.Println(arr1[1:3])

	fmt.Println(&arr1[0],&arr1[1],&arr1[2])
}

func array_2(){
	var arr [3]int32 = [3]int32{2,1,3}
	fmt.Println(arr)

}

func array_3(){
	arr := [3]int32{1,2,3}
	fmt.Println(arr)
}

func array_4(){
	arr := [...]string{"a_s","b_s","ok"}
	fmt.Println(arr)
}

