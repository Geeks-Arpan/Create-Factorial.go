// Time complexity - o(n)
// Space complexity - o(1)

package main
import "fmt"
func factorial(num int) int{
   if num == 1 || num == 0{
      return num
   }
   return num * factorial(num-1)
}
func main(){
   var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	result := factorial(number)
	fmt.Printf("The factorial of %d is: %d\n", number, result)
}