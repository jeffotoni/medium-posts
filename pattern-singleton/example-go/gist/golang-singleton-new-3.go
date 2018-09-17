/*
* Example Pattern Singleton User
* @package     main
* @author      @jeffotoni
* @size        12/09/2018
 */

package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	once.Do(func() {

		instance = make(singleton) // <-- thread safe

	})

	return instance
}

func main() {

	objeto := NewClass() // instance make

	objeto["this"] = "Object NewClass"

	objeto2 := NewClass() // return instance

	fmt.Println("This is:", objeto2["this"])
}
