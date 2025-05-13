package main

import (
	"fmt"
	"os"

	"task-tracker/internal"
)

func main() {
	app, err := internal.NewApplication()
	if err != nil {
		panic(fmt.Sprintf("Error al inicializar la aplicación!! %v", err))
	}

	command, err := app.HandleUserInput()
	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	if command.Action == "add" {
		app.AddTask(command.TaskDescription)

		os.Exit(0)
	}

	if command.Action == "update" {
		app.UpdateTask(command.Id, command.TaskDescription)

		os.Exit(0)
	}
}
