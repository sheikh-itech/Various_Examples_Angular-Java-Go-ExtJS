package main
import"fmt"


func sum(args ... int)int {

	total := 0
	fmt.Print(args," : ")
	for _, num := range args{

		total += num
	}
	return total
}
func main(){

	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3,4))
	arr := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(sum(arr...))

}