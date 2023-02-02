package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"Clarke": "02/03/1978",
		"Steve":  "04/07/1929",
		"Bucky":  "31/12/1985",
	}

	fmt.Println(birthdays)

	//Deleting from a map
	delete(birthdays, "Clarke")
	fmt.Println(birthdays)

	//Slicing in maps
	ages := map[string]int{}
	ages["Clarke"] = 23
	ages["Steve"] = 89
	ages["Bucky"] = 67
	ages["Steve"] = 56 //Over here the older value is going to be overridden and key Steve will be assigned with value 56

	fmt.Println(ages["Steve"])
}
