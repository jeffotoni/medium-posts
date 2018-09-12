/*
* Example Connect
* @package     main
* @author      @jeffotoni
* @size        12/09/2018
 */

package pgconnect

import (
	//"errors"
	"fmt"
	"log"
	"sync"
)

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// call somente
// uma unica vez
var once sync.Once

// Struct driver
type driver struct {
	pgdb *sql.DB
}

var (
	Conn driver
	err  error
)

// connect banco
func (dcon *driver) Connet() *driver {

	once.Do(func() {

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)

		// Struct driver
		dcon.pgdb, err = sql.Open(DB_SORCE, DBINFO) // <--- NOT SAFE

		if err != nil {

			log.Println("Error Connet::", err)
		}
	})

	// em memoria
	return dcon
}

// ping test connexao
func (conn *driver) Ping() string {

	if conn.pgdb.Ping() != nil {

		//return errors.New("Error: ping não funcionou.")
		return "Error: ping não funcionou."
	}

	return "ping"
}

// Get User
func (conn *driver) GetUser(id int) string {

	var email string

	err := conn.pgdb.QueryRow("select email from login where id=$1", id).Scan(&email)

	if err != nil {

		log.Println(err)
		return ""
	}

	return email
}

// Get User
func (conn *driver) AddUser(nome, email string) bool {

	stmt, err := conn.pgdb.Prepare("INSERT INTO login(nome,email)VALUES($1,$2)")

	if err != nil {
		log.Println(err)
		return false
	}

	res, err := stmt.Exec(nome, email)

	if err != nil {
		log.Println(err)
		return false
	}

	_, err = res.RowsAffected()

	if err != nil {

		log.Println(err)
		return false
	}

	return true
}
