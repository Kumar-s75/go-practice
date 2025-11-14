///pass struct as function arguments

package main
import ("fmt")

type person struct{
	name string
	age int
	job string
	salary int
}
func main(){
  var pers1 person
  var pers2 person

  pers1.name="Hege"
  pers1.age=45
  pers1.job="Teacher"
  pers1.salary=6000

  pers2.name="Cecilie"
  pers2.age=24
  pers2.job="Marketing"
  pers2.salary=4500

    // Print Pers1 info by calling a function
  printPerson(pers1)

  // Print Pers2 info by calling a function
  printPerson(pers2)
}

func printPerson(pers Person){
	fmt.Println("Name: ",pers.name)
	fmt.Println("Age: ",pers.age)
	fmt.Println("Job: ",pers.job)
	fmt.Println("Salary:",pers.salary)
}

//go maps
////var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
/////b := map[KeyType]ValueType{key1:value1, key2:value2,...}
///Create Maps Using var and :=
package main
import ("fmt")

func main(){
	var a=map[string]string{"brand":"ford","model":"Mustang","year":1964}
	    b:=map[string]int{"Oslo":1,"Bergen":2,"Trondheim":3,"Stavanger":4}

		fmt.Printf("a\t%v\n",a)
		fmt.Printf("b\t%v\n",b)
}

////Create Maps Using the make() Function:
/////var a = make(map[KeyType]ValueType)
///// b := make(map[KeyType]ValueType)
package main
import ("fmt")

func main(){
   var a=make(map[string]string)
   a["brand"]="Ford"
   a["model"]="Mustang"
   a["year"]="1964"

    b := make(map[string]int)

	 b["Oslo"] = 1
  b["Bergen"] = 2
  b["Trondheim"] = 3
  b["Stavanger"] = 4

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

//create an empty map
// var a map[KeyType]ValueType

package main
import ("fmt")

func main(){
	var a = make(map[string]string)
	var b map [string]string

fmt.Println(a == nil)
fmt.Println(b == nil)
	

}

////access map elements
///update,add and delete elements
package main
import ("fmt")

func main(){
	var a=make(map[string]string)
a["brand"] = "Ford"
a["model"] = "Mustang"
a["year"] = "1964"

fmt.Printf(a["brand"])
fmt.Println(a)

a["year"]="1970"
a["color"]="red"

fmt.Println(a)
delete(a,"year")
fmt.Println(a)
}

////check for specific elements in map
// val, ok :=map_name[key]
package main
import ("fmt")

func main(){
	var a=map[string]string{"brand":"ford","model":"Mustang","year":"1964","day":""}
	 val1, ok1 := a["brand"] // Checking for existing key and its value
  val2, ok2 := a["color"] // Checking for non-existing key and its value
  val3, ok3 := a["day"]   // Checking for existing key and its value
  _, ok4 := a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1,ok1)
	fmt.Println(val2,ok2)
	fmt.Println(val3,ok3)
	fmt.Println(ok4)

}

/////maps are references
package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := a

  fmt.Println(a)
  fmt.Println(b)

  b["year"] = "1970"
  fmt.Println("After change to b:")

  fmt.Println(a)
  fmt.Println(b)
}

////iterate over maps
package main
import("fmt")

func main(){
	a:=map[string]int{"one":1,"two":2,"three":3,"four":4}
	for k,v :=range a{
		fmt.Printf("%v:%v,",k,v)
	}
}

//////iterate over maps in specific order
package main
import ("fmt")

func main(){
  a:=map[string]int{"one":1,"two":2,"three":3,"four":4}
  var b []string
  b=append(b,"one","two","three","four")

  for k,v :=range a{
    fmt.Printf("%v:%v",k,v)
  }
  fmt.Println()

  for _,element :=range b{
    fmt.Printf("%v:%v,",element,a[element])
  }
}