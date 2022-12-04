package main //package declaration
import "fmt" //importing the modules

//struct --> we can create collection of members of diff datatypes into single variable
//A struct can be useful for grouping data together to create records.

/* syntax
type struct_name struct {
	member1 datatype;
	member2 datatype;
	.....
}
*/

//diff types of datatypes
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

/*access the members --> to access any member of a structure  we can use the (.)dot
b/w the structure variable name and the structure member:*/

func main() {
	var p1 Person

	//person specification
	p1.name = "siri"
	p1.age = 24
	p1.job = "IT"
	p1.salary = 29000

	// Print Person info by calling a function
	fmt.Println(p1)
	siri(p1)

}

func siri(val Person) { //we can pass the arguments here
	fmt.Println(val.name)
}

// := short hand declaration
