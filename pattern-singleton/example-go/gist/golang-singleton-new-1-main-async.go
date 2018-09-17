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

	go func() {

		for i := 0; i < 1000; i++ {

			objeto = NewClass() // <-- return instance
		}
	}()

	go func() {

		for i := 0; i < 100; i++ {

			go func() {

				objeto = NewClass() // <-- return instance
			}()
		}
	}()

	fmt.Scanln()
}
