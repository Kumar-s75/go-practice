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

package main
import("fmt")

func main(){
	var i,j string="Hello","World"
	fmt.Print(i)
	fmt.Print(j)
	 fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}

package main 
import ("fmt")

func main(){
	var i string="Hello"
	var j int =15
    
	fmt.Printf("i has value:%v and type:%T\n",i,i);
	fmt.Printf("j has value:%v and type:%T",j,j);
}
//%v is used to print value 
//%t is used to print type 

 fmt.Print(i, "\n")
  fmt.Print(j, "\n")

  ///boolean data type
  package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}

//signed and unsigned integers
package main
import ("fmt")

func main(){
	var x int=500
	var y int=-4500
	var z int=-4500
	var a float32=123.78
	var b float32=3.4e+38
	var c float64=1.7e+308
	fmt.Printf("Type:%T,value:%v",x,x);
	fmt.Printf("Type:%T,value:%v",y,y);
	fmt.Printf("Type:%T,value:%v",z,z);
	fmt.Printf("Type:%T,value:%v",a,a);
	fmt.Printf("Type:%T,value:%v",b,b);
	fmt.Printf("Type:%T,value:%v",c,c);
}
