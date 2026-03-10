package service

import (
	"fmt"
	"time"

	"github.com/sanzhar857/task-tracker/internal/models"
	"github.com/sanzhar857/task-tracker/internal/repository"
)

func Help() {
	fmt.Println(`
		┌─────────────────────────────────────────────────────────────────────────┐
		│ Commands:                                                               │
		│  add <task>              - Adding a new task                            │
		│  update <id> <desc>      - Updating tasks                               │
		│  delete <id>             - Deleting tasks                               │
		│  mark-in-progress <id>   - Marking a task as in progress                │
		│  mark-done <id>          - Marking a task as done                       │
		│  list [status]           - Listing tasks (todo, in-progress, done)      │
		│  list                    - Listing all tasks                            │
		│  exit                    - Exit program                                 │
		└─────────────────────────────────────────────────────────────────────────┘
	`)
}

func AddTask(description string) error {
	tasks, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	now := time.Now()

	task := models.Task{
		Id:          repository.NextId(tasks),
		Description: description,
		Status:      models.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	if err := repository.SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task successfully added: {ID:%d}\n", task.Id)

	return nil
}

func UpdateTask(id int, description string) error {
	tasks, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			if err := repository.SaveTasks(tasks); err != nil {
				return err
			}

			fmt.Printf("Task %d updated successfully\n", id)
			return nil
		}
	}

	return fmt.Errorf("task with ID: %d not found", id)
}

func Delete(id int) error {
	tasks, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := repository.SaveTasks(tasks); err != nil {
				return err
			}
			fmt.Printf("Task ID: %d successfully deleted!\n", id)
			return nil
		}
	}

	return fmt.Errorf("Task ID: %d not found!\n", id)
}