# Go Todo List CLI

This project is a simple CLI application for managing a todo list. It allows you to add, delete, list, and mark tasks as completed. The tasks are stored in a JSON file.

## Installation

1. Clone the repository:
	```sh
	git clone https://github.com/yourusername/go-todo-list.git
	cd go-todo-list
	```

2. Build the application:
	```sh
	go build -o todo
	```

## Usage

### Add a new task

To add a new task, use the `add` command followed by the task description:
```sh
./todo add "Buy groceries"
```

### List all tasks


To list all tasks, use the `list` command:
```sh
./todo list
```

```sh
Title             Completed   CreatedAt                  CompletedAt
Buy groceries     No          2023-10-01T12:00:00Z       
Finish the report No          2023-10-01T13:00:00Z       
```

### Delete a task

To delete a task, use the `delete` command followed by the task index:
```sh
./todo delete 0
```

### Mark a task as completed

To mark a task as completed, use the `completed` command followed by the task index:
```sh
./todo completed 0
```

## Example

Here is an example of how to use the CLI:

1. Add a new task:
	```sh
	./todo add "Finish the report"
	```

2. List all tasks:
	```sh
	./todo list
	
	 -----------------------------------------------------------------------------
	│  Title  │ Completed │         CreatedAt         │        CompletedAt        │
	├─────────┼───────────┼───────────────────────────┼───────────────────────────┤
	│ tarea 1 │ Si        │ 2025-03-03T15:24:46-03:00 │ 2025-03-03T15:24:57-03:00 │
	├─────────┼───────────┼───────────────────────────┼───────────────────────────┤
	│ Tarea 2 │ No        │ 2025-03-03T15:25:08-03:00 │                           │
	├─────────┼───────────┼───────────────────────────┼───────────────────────────┤
	│ Tarea 3 │ No        │ 2025-03-03T15:25:14-03:00 │                           │
	├─────────┼───────────┼───────────────────────────┼───────────────────────────┤
	│ Tarea 3 │ No        │ 2025-03-03T18:49:22-03:00 │                           │
	```

3. Mark the first task as completed:
	```sh
	./todo completed 0
	```

4. Delete the first task:
	```sh
	./todo delete 0
	```

## License

This project is licensed under the MIT License.