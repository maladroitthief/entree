package app

import (
	"github.com/maladroitthief/entree/app/command"
	"github.com/maladroitthief/entree/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Update command.UpdateHandler
}

type Queries struct {
	WindowSettings query.WindowSettingsHandler
}
