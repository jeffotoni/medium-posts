/*
* Example Pattern Singleton User
* @package     main
* @author      @jeffotoni
* @size        12/09/2018
 */

package main

import (
	"fmt"
)

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	if instance == nil {

		instance = make(singleton) // <-- not thread safe
	}

	return instance
}

func main() {

	objeto := NewClass() // instance make

	objeto["this"] = "Object NewClass"

	objeto2 := NewClass() // return instance

	fmt.Println("This is:", objeto2["this"])

	objeto3 := NewClass() // return instance
	objeto4 := NewClass() // return instance
	objeto5 := NewClass() // return instance

	fmt.Println("This is:", objeto3["this"])
	fmt.Println("This is:", objeto4["this"])
	fmt.Println("This is:", objeto5["this"])
}
