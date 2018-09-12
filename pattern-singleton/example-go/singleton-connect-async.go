/*
* Example Pattern Singleton User
* @package     main
* @author      @jeffotoni
* @size        12/09/2018
 */

package main

import (
	"fmt"
	"time"
)
import (
	. "github.com/jeffotoni/medium-posts/pattern-singleton/pgconnect"
)

// adiciona usuario
// na base
func main() {

	// conexao
	Db := Conn.Connet()

	// Exemplo 1
	// posso fazer milhares de
	// chamadas e ele retorna sempre
	// a instancia que esta em memoria
	go func() {

		for i := 0; i < 10000000; i++ {

			Db = Conn.Connet()
			fmt.Println("Goroutine1 Connect: ", i)
			time.Sleep(time.Millisecond * 50)
		}
	}()

	// Exemplo 2
	for x := 0; x < 200; x++ {

		go func(x int) {

			for j := 0; j < 10000; j++ {

				fmt.Println("Goroutine2 Connect: ", x, " -> ", j, " login: ", Conn.Connet().GetUserEmail(x))
				time.Sleep(time.Millisecond * 150)
			}
		}(x)
	}

	fmt.Scanln()
}
