package main

import (
	"code.google.com/p/rsc/qr"
	"database/sql"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fumiyas/qrc/lib"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// generate_otp_uri takes in an email and string and spits out an otp
// uri per the spec:
// https://code.google.com/p/google-authenticator/wiki/KeyUriFormat
// otpauth://TYPE/LABEL?PARAMETERS
//
// Example:
// otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example
func generate_otp_uri(email string, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", email, secret)
}

func display_qr(email string, secret string) {
	text := generate_otp_uri(email, secret)
	code, _ := qr.Encode(text, qr.M)
	fmt.Printf("Code for %s (secret %s):\n", email, secret)
	qrc.PrintAA(os.Stdout, code, false)
	fmt.Println()
}

func dump() {
	fmt.Println("Dumping databse...")
	db, err := sql.Open("sqlite3", "./example.db")
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
		display_qr(email, secret)
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
