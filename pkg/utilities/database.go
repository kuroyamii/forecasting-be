package utilities

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

/*
GetDatabase returns database connected to a programmatically defined database connection.
It not throws an error if connection fails in a Fatal manner
But this function will wait 1 second before retrying to connect to db again
Because database connection is essential to almost any backend
*/
func GetDatabase(dbAddress string, dbUsername string, dbPassword string, dbName string) *sql.DB {
	log.Printf("%v GetDatabase database connection: starting database connection process\n", Info("INFO"))

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		dbUsername, dbPassword, dbAddress, dbName)

	var (
		db  *sql.DB
		err error
	)
	for {
		db, err = sql.Open("mysql", dataSourceName)
		if pingErr := db.Ping(); err != nil || pingErr != nil {
			if err != nil {
				log.Printf("%v GetDatabase sql open connection fatal error: %v\n", Red("ERROR"), err)
			} else if pingErr != nil {
				log.Printf("%v GetDatabase db ping fatal error: %v\n", Red("ERROR"), pingErr)
			}
			log.Printf("%v GetDatabase re-attempting to reconnect to database...\n", Info("INFO"))
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	log.Printf("%v GetDatabase database connection: established successfully with %s\n", Info("INFO"), dataSourceName)
	return db
}
