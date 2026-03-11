# task-cli

Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. This project will help you practice your programming skills, including working with the filesystem, handling user inputs, and building a simple CLI application.

## Installation

```bash
git clone https://github.com/sanzhar857/task-tracker.git
cd task-tracker
cd cmd
go build -o task-cli .
```

## Usage

### Add a task

```bash
./task-cli add "Buy groceries"
# Task added successfully (ID: 1)
```

### Update a task

```bash
./task-cli update 1 "Buy groceries and cook dinner"
# Task 1 updated successfully
```

### Delete a task

```bash
./task-cli delete 1
# Task 1 deleted successfully
```

### Mark a task as in progress

```bash
./task-cli mark-in-progress 1
# Task 1 marked as in-progress
```

### Mark a task as done

```bash
./task-cli mark-done 1
# Task 1 marked as done
```

### List tasks

```bash
# All tasks
./task-cli list

# Only done
./task-cli list done

# Only todo
./task-cli list todo

# Only in progress
./task-cli list in-progress
```