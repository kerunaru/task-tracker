package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Database struct {
	name string
	data map[string]Task
}

func NewDatabase(name string) (*Database, error) {
	_, err := os.Stat(name)
	if err != nil {
		createEmptyJSON(name)
	}

	file, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo JSON de BBDD: %w", err)
	}

	var data map[string]Task
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("Error al procesar el archivo JSON de BBDD: %w", err)
	}

	db := &Database{
		name: name,
		data: data,
	}

	return db, nil
}

func createEmptyJSON(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("Error al crear el archivo JSON: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("{}")
	if err != nil {
		return fmt.Errorf("Error al escribir en el archivo JSON: %w", err)
	}

	return nil
}

func (d *Database) CommitChanges() error {
	dir := filepath.Dir(d.name)
	tempFile, err := os.CreateTemp(dir, "db.tmp")
	if err != nil {
		return fmt.Errorf("No se puede crear el archivo temporal %w", err)
	}

	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	rawData, err := json.Marshal(d.data)
	if err != nil {
		return fmt.Errorf("La base de datos no es un JSON válido! %w", err)
	}
	if _, err := tempFile.Write(rawData); err != nil {
		tempFile.Close()

		return fmt.Errorf("No se ha podido escribir el archivo JSON! %w", err)
	}

	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("No se ha podido cerrar el archivo temporal! %w", err)
	}

	if err := os.Rename(tempPath, d.name); err != nil {
		return fmt.Errorf("No se ha podido renombrar el archivo temporal! %w", err)
	}

	if err := os.Chmod(d.name, os.FileMode(0644)); err != nil {
		return fmt.Errorf("No se ha podido establecer los permisos del archivo JSON! %w", err)
	}

	return nil
}

func (d *Database) GetTask(id string) (*Task, error) {
	value, valueFound := d.data[id]
	if !valueFound {
		return nil, fmt.Errorf("Tarea no encontrada: %s", id)
	}

	return &value, nil
}

func (d *Database) UpdateTask(task *Task) error {
	value, valueFound := d.data[task.Id.String()]
	if !valueFound {
		return fmt.Errorf("Tarea no encontrada: %s", task.Id.String())
	}

	value.UpdatedAt = time.Now()
	d.data[task.Id.String()] = *task

	return nil
}

func (d *Database) DeleteTask(id string) error {
	_, valueFound := d.data[id]
	if !valueFound {
		return fmt.Errorf("Tarea no encontrada: %s", id)
	}

	delete(d.data, id)

	return nil
}

func (d *Database) GetAllTasks(taskStatus string) map[string]Task {
	if taskStatus == "" {
		return d.data
	}

	tasks := make(map[string]Task)
	for _, task := range d.data {
		if task.Status == taskStatus {
			tasks[task.Id.String()] = task
		}
	}

	return tasks
}
