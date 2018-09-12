/*
* Example Pattern Singleton Once
* @package     main
* @author      @jeffotoni
* @size        12/09/2018
 */

package main

import (
	"fmt"
	"os"
	"time"
)
import (
	. "github.com/jeffotoni/medium-posts/pattern-singleton/pgconnect"
)

func main() {

	// fazendo varias chamadas
	// porem ele ira sempre
	// instanciar uma unica vez

	// instancia uma
	// unica vez
	// Db := Conn.Connet()

	for j := 1; j < 100; j++ {

		//fmt.Println(Db.GetUser(j))
		fmt.Println(Conn.Connet().GetUser(j))
	}

	time.Sleep(time.Second * 2)
	os.Exit(0)

	// Exemplo 1
	// posso fazer milhares de
	// chamadas e ele retorna sempre
	// a instancia que esta em memoria
	// for i := 0; i < 1000000; i++ {

	// 	Db, err = Conn.Connet()
	// }

	// // Exemplo 2
	// for x := 0; x < 200; x++ {

	// 	fmt.Println("Goroutine: ", x)

	// 	go func(x int) {

	// 		for j := 0; j < 10000; j++ {

	// 			time.Sleep(time.Millisecond * 150)
	// 			Conn.Connet()
	// 			fmt.Println("Goroutine Connect: ", x, " -> ", j)
	// 		}
	// 	}(x)
	// }

	// // Exemplo 3
	// // retorna o que
	// // encontra-se em
	// // memoria
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()
	// Db, err = Conn.Connet()

	// if err != nil {

	// 	fmt.Println("Ping: error, confira a conexão com seu banco")
	// 	os.Exit(0)
	// }

	// if Db.Ping() == nil {

	// 	fmt.Println("Connect com sucesso")

	// } else {

	// 	fmt.Println("Ping: error, confira a conexão com seu banco!")
	// 	os.Exit(0)
	// }

	fmt.Scanln()
}
