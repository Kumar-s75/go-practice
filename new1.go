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
/////access element of a slice
package main
import  ("fmt")

func main(){
	prices := []int{10,20,30}

	fmt.Println(prices[20])
	fmt.Println(prices[0])
	fmt.Println(prices[2])
}

/////change elements of a slice
package main
import ("fmt")

func main(){
	prices:=[]int{10,20,30}
	prices[2]=50
	fmt.Println(prices[0])
	fmt.Println(prices[1])
}

///append elements to a slice
////slice_name = append(slice_name, element1, element2, ...)
package main
import ("fmt")

func main() {
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
}

///append one slice to another slice
// slice3 = append(slice1, slice2...)
package main
import("fmt")

func main(){
	myslice1 := []int{1,2,3}
	myslice2 := []int{4,5,6}
	myslice3 := append(myslice1,myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
  fmt.Printf("length=%d\n", len(myslice3))
  fmt.Printf("capacity=%d\n", cap(myslice3))

}

//chnage length of a slice
package main
import ("fmt")

func main() {
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  myslice1 := arr1[1:5] // Slice array
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = arr1[1:3] // Change length by re-slicing the array
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
}

//memory efficiency
// copy(dest, src)
package main
import ("fmt")

func main() {
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers)
  fmt.Printf("length = %d\n", len(numbers))
  fmt.Printf("capacity = %d\n", cap(numbers))

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

///////operators
package main
import("fmt")

func main(){
	var a=15+25
	fmt.Println(a)
}
//+ operator
package main()
import("fmt")

func main(){
	var (
		sum1=100+50
		sum2=sum1+250
		sum3=sum2+sum2
	)
	fmt.Println(sum3)
}

//assignment operator
package main
import("fmt")

func main(){
	var x=10
	fmt.Println(x)
}

//addition operator
package main
import ("fmt")

func main() {
  var x = 10
  x +=5
  fmt.Println(x)
}

///comparison operator
package main
import("fmt")

func main(){
	var x=5
	var y=10
fmt.Println(x>y)
}
///conditional statements
///if statement
package main
import ("fmt")

func main(){
	if 20>18{
		fmt.Println("20 is greater than 18");
	}
	x:=20
	y:=18
	if x>y{
		fmt.Println("x is greater than y");
	}
}

///if else statement
package Main
import ("fmt")

func main(){
	time:=20
	if(time>18){
		fmt.Println("Good day");

	}else{
		fmt.Println("Good evening");
	}

}
////nested if
package main
import ("fmt")

func main(){
	num:=20
	if num>=10{
		fmt.Println("Num is more than 10.");
	if num >=16{
		fmt.Println("Num is more than 16");
	} else{
		fmt.Println("Num is less than 10");
	}
	}
}

///switch case
package main
import ("fmt")

func main(){
	day:=4
	switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}

///multi case switch example

package main
import ("fmt")
 
func main(){
	day:=5

	switch day{
	case 1,3,5:
		fmt.Println("Odd weekday")
    case 2,4:
		fmt.Println("Even weekday")
	case 6,7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

//loop examples
//for loop
package main
import ("fmt")

func main(){
	for i:=0;i<5;i++{
		fmt.Println(i)
	}
}
///continue statement
package main
import ("fmt")

func main(){
	for i:=0;i<5;i++{
		if i==3{
			continue
		}
		fmt.Println(i)
	}
}
//break statement
package main
import ("fmt")


func main(){
 for i:=0;i<5;i++{
	if i==3{
		break
	}
	fmt.Println(i)
 }
}

//nested loops
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}
///range keyword
package main
import ("fmt")
func main(){
	fruits:=[3]string{"apple","orange","banana"}
	for idx,val:=range fruits{
		fmt.Printf("%v\t%v\n",idx,val)
	}
}

///function
package main
import ("fmt")

func myMessage(){
	fmt.Println("I just got executed!")

}
func main(){
	myMessage();
	myMessage();
	myMessage();
}

package main
import ("fmt")
///single parameter
func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}
/////multiple parameter
package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}

////recursion functions
package main
import ("fmt")

func testcount(x int){
 if x==11{
	return 0
 }
 fmt.Println(x)
 return testcount(x+1)

}
func main(){
	testcount(1)
}
////struct
type person struct{
	name string
	age int
	job string
	salary int
}
////access struct member
package main
import("fmt")

type Person struct{
	name string
	age int
	job string
	salary int

}

func main(){
	var pers1 Person
	var pers2 Person

	pers1.name="Hege"
	pers1.age=45
	pers1.job="Teacher"
	pers1.salary=6000

	 // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

  // Access and print Pers1 info
  fmt.Println("Name: ", pers1.name)
  fmt.Println("Age: ", pers1.age)
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)

  // Access and print Pers2 info
  fmt.Println("Name: ", pers2.name)
  fmt.Println("Age: ", pers2.age)
  fmt.Println("Job: ", pers2.job)
  fmt.Println("Salary: ", pers2.salary)
}
