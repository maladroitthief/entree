package app

import "github.com/maladroitthief/entree/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	Update command.UpdateHandler
}
