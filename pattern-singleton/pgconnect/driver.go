/*
* Example Connect
* @package     main
* @author      @jeffotoni
* @size        12/09/2018
 */

package pgconnect

import (
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

func (dcon *driver) Connet() (*sql.DB, error) {

	once.Do(func() {

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)

		// Struct driver
		dcon.pgdb, err = sql.Open(DB_SORCE, DBINFO) // <--- NOT SAFE

		if err != nil {

			log.Println("Error Connet::", err)
		}
	})

	return dcon.pgdb, err
}
