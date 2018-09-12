/*
* Example Pattern Singleton User
* @package     main
* @author      @jeffotoni
* @size        12/09/2018
 */

package main

import (
	"encoding/json"
	"fmt"
	"time"
)
import (
	. "github.com/jeffotoni/medium-posts/pattern-singleton/pgconnect"
)

type login struct {
	Nome  string
	Email string
}

// adiciona usuario
// na base
func main() {

	var channel = make(chan string)

	var j int

	go func(j int, channel chan string) {

		idunico := 10

		//L := &login{}

		for j = 0; j < 100; j++ {

			nome := "arthur " + fmt.Sprintf("%d", idunico) + " " + fmt.Sprintf("%d", j)
			email := "arthur" + fmt.Sprintf("%d", idunico) + "-" + fmt.Sprintf("%d", j) + "@gmail.com"

			bjson, err := json.Marshal(&login{nome, email})
			if err != nil {
				fmt.Println("error marshal:", err)
			}

			channel <- string(bjson)

		}
	}(j, channel)

	go func(j int, channel chan string) {

		idunico := 2

		for j = 0; j < 100; j++ {

			nome := "francisco " + fmt.Sprintf("%d", idunico) + " " + fmt.Sprintf("%d", j)
			email := "francisco" + fmt.Sprintf("%d", idunico) + "-" + fmt.Sprintf("%d", j) + "@gmail.com"

			bjson, err := json.Marshal(&login{nome, email})
			if err != nil {
				fmt.Println("error marshal:", err)
			}

			channel <- string(bjson)
		}
	}(j, channel)

	go func(channel chan string) {

		L := login{}

		for cjson := range channel {

			if cjson != "" {

				fmt.Println(cjson)
				err := json.Unmarshal([]byte(cjson), &L)

				if err != nil {
					fmt.Println("error:", err)
				}

				if Conn.Connet().AddUser(L.Nome, L.Email) {

					fmt.Println("Adicionado: User", L.Nome, " Email: ", L.Email)
				}

				time.Sleep(time.Second * 1)
			}
		}
	}(channel)

	//time.Sleep(time.Second * 20)
	fmt.Scanln()
}
