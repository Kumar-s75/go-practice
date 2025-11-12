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
	var txt1 string="Hello!"
	var txt2 string
	txt3:="World 1"

	fmt.Printf("Type:%T,value:%v",x,x);
	fmt.Printf("Type:%T,value:%v",y,y);
	fmt.Printf("Type:%T,value:%v",z,z);
	fmt.Printf("Type:%T,value:%v",a,a);
	fmt.Printf("Type:%T,value:%v",b,b);
	fmt.Printf("Type:%T,value:%v",c,c);
	fmt.Printf("Type:%T,value:%v\n",txt1,txt1);
	fmt.Printf("Type:%T,value:%v\n",txt2,txt2);
	fmt.Printf("Type:%T,value:%v\n",txt3,txt3);
}

////string
package main
import ("fmt")
func main(){
	var txt1 string="Hello!"
	var txt2 string
	txt3:="World 1"

	fmt.Printf("Type:%T,value:%v\n",txt1,txt1);
	fmt.Printf("Type:%T,value:%v\n",txt2,txt2);
	fmt.Printf("Type:%T,value:%v\n",txt3,txt3);
}
//Arrays
var array_name=[Length]datatypes{values}
or
var array_name=[...]datatype{values}

array_name:=[Length]datatype{values}
array_name:=[...]datatype{values}
////[x] defined
////[...] inferred

package main
import ("fmt")

func main(){
	var arr1=[3]int{1,2,3}
	arr2:=[5]int{4,5,6,7,8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
//////
package main
import ("fmt")

func main(){
	var arr1=[...]int{1,2,3}
	arr2:=[...]int{4,5,6,7,8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
///////
package main
import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}

/////slices part1
/////slice_name := []datatype{values}
package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}

///slices part2
///create a slice from an array
// var myarray = [length]datatype{values} // An array
// myslice := myarray[start:end] // A slice made from the array
package main
import ("fmt")

func main(){
	arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
}

///creating a slice from the make function

////slice_name := make([]type, length, capacity)
package main
import("fmt")

func main(){
	myslice1:=make([]int,5,10)
	fmt.Printf("myslice1=%v\n",myslice1);
	fmt.Printf("length=%d\n",len(myslice1));
	fmt.Printf("capacity=%d\n",cap(myslice1));

	//with omitted capacity
	myslice2:=make([]int,5);
	fmt.Printf("myslice2=%v\n",myslice2);
	fmt.Printf("length=%d\n",cap(myslice1));
	fmt.Printf("capacity = %d\n", cap(myslice2));

}