package main

import (
	"fmt"
)

type Student struct {
	name       string;
	id         int;
	GPA        float32;
	department string;
}

func main() {
	fmt.Println("Value: goo");
	const b = iota;
	const c = iota;
	const d = iota;
	const a complex64= 1 + 2i;
	//fmt.Printf("%v, %v, %v", a, c, d);
	var zero Student = Student { name: "Cicero", id: 0002, GPA: 3.9, department: "Liberal Arts",};
	zero.GPA = zero.GPA;
	fmt.Printf("%v of the %v department has a GPA of %v.", zero.name, zero.department, zero.GPA);


}
