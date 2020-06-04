package main

import (
	"fmt"

	"github.com/braulioinf/QRLauncher/src/app"
	"github.com/gorilla/mux"
)

func main() {

	app := app.App{
		Router: mux.NewRouter().StrictSlash(true),
	}

	if err := app.Initialize(); err != nil {
		fmt.Printf("MySQL error conexion. %v\n", err)
	}

	app.SetupRouter()
	app.Run()
}
