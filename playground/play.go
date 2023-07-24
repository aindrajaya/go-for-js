package playground
import "fmt"

nestedArray := \[3\] [5]int{
	{1,2,3,4,5},
	{6,7,8,9,10},
	{11,12,13,14,15},
}

func play() {
	fmt.Println(nestedArray)
}
