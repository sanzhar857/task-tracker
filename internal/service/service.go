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

func MarkInProgress(id int) error {
	tasks, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = models.StatusInProgress
			tasks[i].UpdatedAt = time.Now()
			if err := repository.SaveTasks(tasks); err != nil {
				return err
			}

			fmt.Printf("Task ID: %d successfully marked to in-progress\n", id)
			return nil
		}
	}
	return fmt.Errorf("Task ID: %d not found!\n", id)
}

func MarkDone(id int) error {
	tasks, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = models.StatusDone
			tasks[i].UpdatedAt = time.Now()
			if err := repository.SaveTasks(tasks); err != nil {
				return err
			}
			fmt.Printf("Task ID %d successfully marked to Done\n", id)
			return nil
		}
	}
	return fmt.Errorf("Task ID: %d not found!\n", id)
}

func ListTasks(filter string) error {
	tasks, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	filtered := []models.Task{}

	for _, t := range tasks {
		switch filter {
		case "todo":
			if t.Status == models.StatusToDo {
				filtered = append(filtered, t)
			}
		case "in-progress":
			if t.Status == models.StatusInProgress {
				filtered = append(filtered, t)
			}
		case "done":
			if t.Status == models.StatusDone {
				filtered = append(filtered, t)
			}
		default:
			filtered = append(filtered, t)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("Not tasks found")
		return nil
	}

	print_tasks(filtered)
	return nil
}

func print_tasks(tasks []models.Task) {
	fmt.Printf("%-4s | %-30s | %-15s | %-20s | %-20s\n", "ID", "Description", "Status", "Created Time", "Updated Time")
	fmt.Println("-----------------------------------------------------------------------------------------------------")

	for _, task := range tasks {
		fmt.Printf("%-4d | %-30s | %-15s | %-20v | %-20v\n", task.Id, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}
