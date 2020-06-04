package app

import (
	"github.com/braulioinf/QRLauncher/src/controllers/content"
)

// SetupRouter func
func (a *App) SetupRouter() {
	CtrlCont := content.NewCtrl(a.Database)

	a.Router.Methods("POST").Path("/api/contents").HandlerFunc(CtrlCont.Create)

}
