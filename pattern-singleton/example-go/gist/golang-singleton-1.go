/*
* Example DriverPg Go
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
 */

package main

import (
	"fmt"
	"time"
)

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {

	if instance == nil {
		// <--- NOT THREAD SAFE / Quando usarmos Goroutine
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}

	return instance
}

func main() {

	// chamada
	go func() {

		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 600)
			fmt.Println(*Connect(), " - ", i)
		}
	}()

	go func() {

		fmt.Println(*Connect())
	}()

	fmt.Scanln()
}
