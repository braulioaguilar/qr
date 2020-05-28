package app

import (
	"database/sql"
	"fmt"
)

type App struct {
	Database *sql.DB
}

func (a *App) Run() {
	fmt.Println("Running...")
}
