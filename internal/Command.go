package internal

import (
	"errors"
	"strconv"
)

type Command struct {
	Action          string
	Id              int
	TaskDescription string
	TaskStatus      string
}

func NewAddCommand(action string, taskDescription string) (*Command, error) {
	if action != "add" {
		return nil, errors.New("Command action must be \"add\"!")
	}

	if taskDescription == "" {
		return nil, errors.New("Task description must not be empty!")
	}

	command := &Command{
		Action:          action,
		TaskDescription: taskDescription,
	}

	return command, nil
}

func NewUpdateCommand(action string, rawId string, taskDescription string) (*Command, error) {
	if action != "update" {
		return nil, errors.New("Command action must be \"update\"!")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, err
	}

	if id < 1 {
		return nil, errors.New("Task ID must be greater than 0!")
	}

	if taskDescription == "" {
		return nil, errors.New("Task description must not be empty!")
	}

	command := &Command{
		Action:          action,
		Id:              id,
		TaskDescription: taskDescription,
	}

	return command, nil
}

func NewDeleteCommand(action string, rawId string) (*Command, error) {
	if action != "delete" {
		return nil, errors.New("Command action must be \"delete\"!")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, err
	}

	if id < 1 {
		return nil, errors.New("Task ID must be greater than 0!")
	}

	command := &Command{
		Action: action,
		Id:     id,
	}

	return command, nil
}

func NewListCommand(action string, taskStatus string) (*Command, error) {
	if action != "list" {
		return nil, errors.New("Command action must be \"list\"!")
	}

	if !(taskStatus == "in-progress" || taskStatus == "done" || taskStatus == "todo") {
		return nil, errors.New("Task status must be one of: in-progress, done or todo!")
	}

	command := &Command{
		Action:     action,
		TaskStatus: taskStatus,
	}

	return command, nil
}
