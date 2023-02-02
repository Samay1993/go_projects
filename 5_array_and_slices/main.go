package main

import "fmt"

func main() {
	names := [3]string{"Adam", "Clarke", "Bruce"} //shorthand to define a variable in this case array (:=)
	fmt.Println(names)

	names[0] = "Barry"
	names[1] = "Tony"
	names[2] = "Peter"

	fmt.Println(names)

	//Slicing of Array
	names_new := []string{}
	names_new = append(names_new, "David")
	names_new = append(names_new, "Silvia", "Greg", "Tom")

	fmt.Println(names_new)
	fmt.Println("names_new[2] is nil: ", names_new[2] == "")

	//Slicing when you know the starting length of the array(in this case 4), after that it will increase length as we go
	name := make([]string, 4)
	name[0] = "Wanda"
	name[1] = "Kent"
	name[2] = "Natasha"
	name[3] = "Pepper"

	fmt.Println(name)

	//Adding more indices and values to the array
	name = append(name, "Happy", "Harry")
	fmt.Println(name)
}
