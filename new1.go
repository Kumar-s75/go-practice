// package main
// import ("fmt")

// func main()
// {
//   fmt.Println("Hello World!")
// }

// var variablename type = value

// variablename := value
// package main
// import ("fmt")
// func Main(){
// 	var student1 string="john"
// 	var student2 ="jae"
// 	x:=2
// 	fmt.Println(student1);
// 	fmt.Println(student2);
// 	fmt.Println(x);
// }
//  Variable Declaration Without Initial Value
package main
import ("fmt")

func Main(){
	var a string
	var b int
	var c bool
 fmt.Println(a)
 fmt.Println(b)
 fmt.Println(c)
}

// value assignment after declaration
package Main
import ("fmt")

func main(){
	var student1 string
	student1="John"
	fmt.Println(student1)
}

//example using both type declaration of variable
package Main
import ("fmt")

var a int
var b int=2
var c=3

func main(){
	a=1
	d:=6
fmt.Println(a)
fmt.println(b)
fmt.println(c)
}

//////
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}