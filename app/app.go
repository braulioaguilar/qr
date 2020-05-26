package app

import (
	"database/sql"
)

type App struct {
	Database *sql.DB
}

func (a *App) Run() {
}
