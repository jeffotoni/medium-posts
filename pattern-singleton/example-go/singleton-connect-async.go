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

	// busca sempre
	// da memoria
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()

	fmt.Println("connect", Db.Ping())
	fmt.Println("connect", Conn.Connet().Ping())
	fmt.Println("connect", Db.Ping())
	fmt.Println("connect", Conn.Connet().Ping())
	fmt.Println("connect", Db.Ping())
	fmt.Println("connect", Conn.Connet().Ping())

	time.Sleep(time.Second * 2)

	// Exemplo 1
	// posso fazer milhares de
	// chamadas e ele retorna sempre
	// a instancia que esta em memoria
	for i := 0; i < 1000000; i++ {

		Db = Conn.Connet()
	}

	// Exemplo 2
	for x := 0; x < 100; x++ {

		//fmt.Println("Goroutine: ", x)

		go func(x int) {

			for j := 0; j < 10000; j++ {

				fmt.Println("Goroutine Connect: ", x, " -> ", j, " login: ", Conn.Connet().GetUserEmail(x))
				time.Sleep(time.Millisecond * 150)
			}
		}(x)
	}

	fmt.Scanln()
}
