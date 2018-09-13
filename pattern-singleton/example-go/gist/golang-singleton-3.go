/*
* Example DriverPg Go
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

// variavel Global
var instance *DriverPg

// funcao retornando
// o ponteiro de nossa
// struct
func Connect() *DriverPg {

	if instance == nil { // ainda não está perfeita, não é totalmente atomico

		lock.Lock()
		defer lock.Unlock()
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}

	return instance
}

func main() {

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
