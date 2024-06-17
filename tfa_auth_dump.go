package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mgutz/ansi"
	"github.com/qpliu/qrencode-go/qrencode"
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

// Actually print the right QR code to stdout given and email and secret
func display_qr(email string, secret string) {
	text := generate_otp_uri(email, secret)
	code, _ := qrencode.Encode(text, qrencode.ECLevelL)
	fmt.Printf("Code for %s (secret %s):\n", email, secret)
	printAA(os.Stdout, code, false)
	fmt.Println()
}

// Dumps and calls out to display qr codes based on a given input file
func dump(db_file string) {
	fmt.Printf("Dumping databse %s...", db_file)
	db, err := sql.Open("sqlite3", db_file)
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

// Vendored from https://raw.githubusercontent.com/fumiyas/qrc/master/lib/aa.go
func printAA(w_in io.Writer, grid *qrencode.BitGrid, inverse bool) {
	// Buffering required for Windows (go-colorable) support
	w := bufio.NewWriterSize(w_in, 1024)

	reset := ansi.ColorCode("reset")
	black := ansi.ColorCode(":black")
	white := ansi.ColorCode(":white")
	if inverse {
		black, white = white, black
	}

	height := grid.Height()
	width := grid.Width()
	line := white + fmt.Sprintf("%*s", width*2+2, "") + reset + "\n"

	fmt.Fprint(w, line)
	for y := 0; y < height; y++ {
		fmt.Fprint(w, white, " ")
		color_prev := white
		for x := 0; x < width; x++ {
			if grid.Get(x, y) {
				if color_prev != black {
					fmt.Fprint(w, black)
					color_prev = black
				}
			} else {
				if color_prev != white {
					fmt.Fprint(w, white)
					color_prev = white
				}
			}
			fmt.Fprint(w, "  ")
		}
		fmt.Fprint(w, white, " ", reset, "\n")
		w.Flush()
	}
	fmt.Fprint(w, line)
	w.Flush()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a sqlite filename")
		return
	}
	file := os.Args[1]
	dump(file)
}
