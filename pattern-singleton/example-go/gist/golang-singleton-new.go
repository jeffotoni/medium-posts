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

// type Singleton struct {
// 	This map[string]string
// }

type singleton map[string]string

var (
	once sync.Once

	instance singleton
)

func NewClass() singleton {

	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

func main() {

	objeto := NewClass() // instance make

	objeto["this"] = "that"

	objeto2 := NewClass() // return instance

	objeto2["this"] = "what here"

	fmt.Println("This is ", objeto2["this"])

	objeto3 := NewClass() // return instance
	objeto4 := NewClass() // return instance
	objeto5 := NewClass() // return instance
	objeto6 := NewClass() // return instance

}
