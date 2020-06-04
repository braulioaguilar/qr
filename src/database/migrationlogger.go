package database

import "log"

// MigrationLogger struct
type MigrationLogger struct {
	verbose bool
}

// Printf func
func (ml *MigrationLogger) Printf(format string, v ...interface{}) {
	log.Printf(format+"\n", v)
}

// Verbose func
func (ml *MigrationLogger) Verbose() bool {
	return ml.verbose
}
