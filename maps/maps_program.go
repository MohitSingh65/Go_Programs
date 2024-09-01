package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Berylium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	var element string
	fmt.Println("Enter key:")
	fmt.Scanln(&element)
	//	fmt.Printf("Element for key %s is %s", element, elements[element])

	if val, ok := elements[element]; ok {
		fmt.Printf("Element for key %s is %s\n", element, val)
	} else {
		fmt.Printf("Element for key %s not found\n", element)
	}
}
