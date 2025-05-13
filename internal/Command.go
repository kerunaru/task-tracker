package internal

import (
	"errors"
)

type Command struct {
	Action          string
	Id              string
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

	if rawId == "" {
		return nil, errors.New("Task ID must not be empty!")
	}

	if taskDescription == "" {
		return nil, errors.New("Task description must not be empty!")
	}

	command := &Command{
		Action:          action,
		Id:              rawId,
		TaskDescription: taskDescription,
	}

	return command, nil
}

func NewDeleteCommand(action string, rawId string) (*Command, error) {
	if action != "delete" {
		return nil, errors.New("Command action must be \"delete\"!")
	}

	if rawId == "" {
		return nil, errors.New("Task ID must be greater than 0!")
	}

	command := &Command{
		Action: action,
		Id:     rawId,
	}

	return command, nil
}

func NewListCommand(action string, taskStatus string) (*Command, error) {
	if action != "list" {
		return nil, errors.New("Command action must be \"list\"!")
	}

	if !(taskStatus == "" || taskStatus == "in-progress" || taskStatus == "done" || taskStatus == "todo") {
		return nil, errors.New("Task status must be empty or one of: in-progress, done or todo!")
	}

	command := &Command{
		Action:     action,
		TaskStatus: taskStatus,
	}

	return command, nil
}

func NewMarkCommand(action string, taskStatus string, id string) (*Command, error) {
	if action != "mark" {
		return nil, errors.New("Command action must be \"list\"!")
	}

	if !(taskStatus == "in-progress" || taskStatus == "done" || taskStatus == "todo") {
		return nil, errors.New("Task status must be one of: in-progress, done or todo!")
	}

	if id == "" {
		return nil, errors.New("Task ID must be greater than 0!")
	}

	command := &Command{
		Action:     action,
		Id:         id,
		TaskStatus: taskStatus,
	}

	return command, nil
}
