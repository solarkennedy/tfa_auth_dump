package main

import (
	"database/sql"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	//	"github.com/fumiyas/qrc/cmd/qrc"
	"fmt"
	"log"
	"os"
)

func dump() {
	fmt.Println("Dumping databse...")
	db, err := sql.Open("sqlite3", "./test.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select email, secret from accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		var secret string
		rows.Scan(&email, &secret)
		fmt.Println(email, secret)
	}
	rows.Close()
}

func main() {

	app := cli.NewApp()
	app.Name = "qr_auth_dump"
	app.Usage = "Dump secrets from a Google Authenticator database and spit out QR codes"
	app.Action = func(c *cli.Context) {
		dump()
	}

	app.Run(os.Args)

}
