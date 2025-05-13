package internal

import (
	"errors"
	"fmt"
	"os"
)

type Application struct {
	DB *Database
}

func NewApplication() (*Application, error) {
	app := &Application{}
	var err error

	app.DB, err = NewDatabase("db.json")
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) HandleUserInput() (*Command, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return nil, errors.New("Uso: task-tracker [add|update|delete|list|mark-in-progress|mark-done]")
	}

	var command *Command
	var err error

	if a.getArgOrDefault(args, 0) == "add" {
		command, err = NewAddCommand(
			a.getArgOrDefault(args, 0),
			a.getArgOrDefault(args, 1),
		)
	}
	if err != nil {
		return nil, fmt.Errorf("Error procesando la entrada del usuario: %w", err)
	}

	return command, nil
}

func (a *Application) AddTask(description string) error {
	newRecord, err := NewTask(description)
	if err != nil {
		return fmt.Errorf("Error al instanciar la tarea: %w", err)
	}

	a.DB.data[len(a.DB.data)] = *newRecord

	if err := a.DB.CommitChanges(); err != nil {
		return fmt.Errorf("Error al persistir los cambios en la base de datos: %w", err)
	}

	return nil
}

func (a *Application) getArgOrDefault(args []string, index int) string {
	if index >= 0 && index < len(args) {
		return args[index]
	}

	return ""
}
