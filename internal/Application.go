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

	if a.getArgOrEmpty(args, 0) == "add" {
		command, err = NewAddCommand(
			a.getArgOrEmpty(args, 0),
			a.getArgOrEmpty(args, 1),
		)
	}

	if a.getArgOrEmpty(args, 0) == "update" {
		command, err = NewUpdateCommand(
			a.getArgOrEmpty(args, 0),
			a.getArgOrEmpty(args, 1),
			a.getArgOrEmpty(args, 2),
		)
	}

	if a.getArgOrEmpty(args, 0) == "delete" {
		command, err = NewDeleteCommand(
			a.getArgOrEmpty(args, 0),
			a.getArgOrEmpty(args, 1),
		)
	}

	if a.getArgOrEmpty(args, 0) == "list" {
		command, err = NewListCommand(
			a.getArgOrEmpty(args, 0),
			a.getArgOrEmpty(args, 1),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("Error procesando la entrada del usuario: %w", err)
	}

	return command, nil
}

func (a *Application) getArgOrEmpty(args []string, index int) string {
	if index >= 0 && index < len(args) {
		return args[index]
	}

	return ""
}

func (a *Application) AddTask(description string) error {
	newRecord, err := NewTask(description)
	if err != nil {
		return fmt.Errorf("Error al instanciar la tarea: %w", err)
	}

	a.DB.data[newRecord.Id.String()] = *newRecord

	if err := a.DB.CommitChanges(); err != nil {
		return fmt.Errorf("Error al persistir los cambios en la base de datos: %w", err)
	}

	return nil
}

func (a *Application) UpdateTask(id string, description string) error {
	task, err := a.DB.GetTask(id)
	if err != nil {
		return fmt.Errorf("Error al actualizar la tarea desde BBDD: %w", err)
	}

	task.Description = description

	a.DB.UpdateTask(task)

	if err := a.DB.CommitChanges(); err != nil {
		return fmt.Errorf("Error al persistir los cambios en la base de datos: %w", err)
	}

	return nil
}

func (a *Application) DeleteTask(id string) error {
	err := a.DB.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("Error al eliminar la tarea desde BBDD: %w", err)
	}

	if err := a.DB.CommitChanges(); err != nil {
		return fmt.Errorf("Error al persistir los cambios en la base de datos: %w", err)
	}

	return nil
}

func (a *Application) ListTasks(taskStatus string) {
	tasks := a.DB.GetAllTasks(taskStatus)
	for _, task := range tasks {
		fmt.Println(fmt.Sprintf("%s %s %s", task.Id.String(), task.Status, task.Description))
	}
}
