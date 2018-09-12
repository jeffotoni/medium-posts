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
		fmt.Println(Conn.Connet().GetUserEmail(j))
	}

	time.Sleep(time.Second * 2)
}
