package main

import (
	"github.com/codegangsta/cli"
//	"github.com/fumiyas/qrc/cmd/qrc"
	"os"
)

func main() {

  app := cli.NewApp()
  app.Name = "qr_auth_dump"
  app.Usage = "Usage string"
  app.Action = func(c *cli.Context) {
    println("Hello main")
  }

  app.Run(os.Args)

}
