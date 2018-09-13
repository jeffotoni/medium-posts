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

// nosso lock mutex
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
	// <--- Desnecessario a lock
	// se a instancia já tiver
	// sido criada muito agressivo
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		// ainda não é a melhor implementação devido
		// os bloqueios
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
