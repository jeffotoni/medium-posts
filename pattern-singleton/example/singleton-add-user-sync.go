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

	// 100 deles
	var j int

	for j = 0; j < 100; j++ {

		nome := "arthur-" + fmt.Sprintf("%d", j)
		email := "arthur" + fmt.Sprintf("%d", j) + "@gmail.com"

		//if !Db.AddUser(nome, email) {
		if !Conn.Connet().AddUser(nome, email) {

			fmt.Println("Erro ao criar usuario")
		}
	}

	fmt.Println("Adicionado", j, " ususarios com sucesso")
	time.Sleep(time.Second * 1)
	//fmt.Scanln()
}
