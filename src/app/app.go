package app

import (
	"database/sql"

	c "github.com/braulioinf/QRLauncher/src/constants"
	"github.com/braulioinf/QRLauncher/src/cors"
	"github.com/braulioinf/QRLauncher/src/database"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type App struct {
	Database *sql.DB
	Router   *mux.Router
}

// Initialize func & db conexion
func (a *App) Initialize() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	a.Database = db

	return nil
}

// Run func
func (a *App) Run() {
	// It is a middleware-focused library that is designed to work directly with net/http.
	n := negroni.Classic()

	// This middleware catches panics and responds with a 500 response code.
	n.Use(negroni.NewRecovery())

	n.Use(negroni.HandlerFunc(cors.SetupOptions))
	n.Use(cors.SetupHeaders())

	n.UseHandler(a.Router)
	n.Run(":" + c.APP_PORT)
}
