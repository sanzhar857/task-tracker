package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/sanzhar857/task-tracker/internal/models"
)

const storageFile = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(storageFile)
	if errors.Is(err, os.ErrNotExist) {
		return []models.Task{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to read file %w\n", err)
	}

	var tasks []models.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to Parse JSON %w\n", err)
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("failed to serialize tasks %w\n", err)
	}

	return os.WriteFile(storageFile, data, 0644)
}

func NextId(tasks []models.Task) int {
	max := 0

	for _, t := range tasks {
		if t.Id > max {
			max = t.Id
		}
	}

	return max + 1
}
