package main

import (
	"web/api"
	"web/models"
)

type App struct {
	router api.Router
	model  models.Model
}

func (a *App) setup() {
	a.model.Setup()
	a.router.Setup()
}

func (a *App) run() {
	a.router.Run()
}

func (a *App) close() {
	a.router.Close()
	a.model.Close()
}

func main() {
	app := &App{}
	app.setup()
	defer app.close()
	app.run()
}
