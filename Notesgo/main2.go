/*	1. Unsigned_int ::= uint , uint8,uint16,uint16,uint32,uint64
	2. int ::= int, int8,int16,int32,int64
	3. float -> cannot be mentioned ideal utilize floadt8 or upscaled versions 
	4. Strings numlen pack*ge import "unicode/utf8"
 */

package main
import "fmt"

func main(){
	fmt.Println("Hello world")

	var intnum int = 194840
	var unintnum uint8 = 879
	fmt.Println(unintnum)

	var floatnum1 float64 = 837.4  
    fmt.Println(floatnum1)


// Cast the int to float 
	var num1_32b float32 = 22.1 
	var num2_32b int32  = 2 
	var res float32 = num1_32b + num2_32b
	fmt.Println(res)
	fmt.Println(num1_32b/num2_32b)
	fmt.Println(num1_32b%num2_32b)

// String types num of bytes in string @go

	var string1 string = `hello world`
	fmt.Println(string1 + "Gnanesh" ) 

	fmt.Println(len("A"))
	fmt.Println(len("Y"))

	// from string pckg not used widely 
	fmt.Println(utf8.RuneCountInString("y"))

// rune in go 
	var runevar1 rune = 'a'
	fmt.Println(runevar1)


// %%%%%%%%%%%%%%%%%%%%%%%%%%ADVANCED utilization%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

 	a1 := "text-any"
 	fmt.Println(a1)

	n1,n2 := 1,2
	fmt.Println(n1,n2)


//constant 
	const c1 string = "const value"
	fmt.Println(c1)

	const pi float32 = 3.14
	fmt.Println(pi)

}

