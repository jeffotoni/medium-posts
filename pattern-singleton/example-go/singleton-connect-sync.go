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

	time.Sleep(time.Second * 1)
}
