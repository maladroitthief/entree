package application

import (
	"github.com/maladroitthief/entree/application/command"
	"github.com/maladroitthief/entree/application/query"
)

type Game struct {
	Commands GameCommands
	Queries  GameQueries
}

type GameCommands struct {
	Update command.UpdateGameHandler
}

type GameQueries struct {
	WindowSettings query.WindowSettingsHandler
}
