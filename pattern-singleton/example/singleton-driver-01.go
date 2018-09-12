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
	Db, err := Conn.Connet()

	// retorna o que
	// encontra-se em
	// memoria
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()
	Db, err = Conn.Connet()

	if err != nil {

		fmt.Println("Ping: error, confira a conexão com seu banco")
		os.Exit(0)
	}

	if Db.Ping() == nil {

		fmt.Println("Connect com sucesso")

	} else {

		fmt.Println("Ping: error, confira a conexão com seu banco!")
		os.Exit(0)
	}
}
