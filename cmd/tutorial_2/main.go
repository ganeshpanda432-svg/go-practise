package main
import "fmt"

func main(){
	var intNum int =2342
	fmt.Println(intNum)

	var floatNum float64=123434.5
	fmt.Println(floatNum)

	var floatNum32 float32=12.3
	var intnum int=23
	var result float32=floatNum32 + float32(intnum)
	fmt.Println(result)

	var intNum1 int=10
	var intNum2 int=4
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var mystring string="Hello "
	fmt.Println(mystring)
}
