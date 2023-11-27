package main

import (
	"log"
	"os"

	_ "github.com/krekka/cit-server-backend/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: true,
	})

	// Importing default collection
	// app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
	// 	form := forms.NewCollectionsImport(app)
	// 	form.Submit()

	// 	return nil
	// })

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
