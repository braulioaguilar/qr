package main

import (
	"fmt"

	"github.com/braulioinf/QRLauncher/src/app"
)

func main() {

	app := app.App{}

	if err := app.Initialize(); err != nil {
		fmt.Printf("MySQL error conexion. %v\n", err)
	}

	app.Run()
}
