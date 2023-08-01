Sure, here is an example README for your `tasker` CLI app:

# Tasker CLI App

`tasker` is a command-line app that helps you manage your tasks. It allows you to create, list, mark as complete, and delete tasks. The app stores tasks in a mongodb database.

## Installation

To install `tasker`, you can download the latest release binary from the [GitHub releases page](https://github.com/irononet/tasker/releases). Once you have downloaded the binary, you can move it to a directory that's included in your system's `PATH` environment variable, such as `/usr/local/bin` on Unix-based systems.

Alternatively, you can clone this repository and build the binary yourself by running `go build -o tasker`.

You can also build  a docker container with all the dependencies by running 

```bash
sudo docker-compose up
```

## Usage

Here are the available commands in `tasker`:

### `tasker add (a)`

Add a new task to the list. You will be prompted for the task description.

Example:

```
tasker add
```

### `tasker all (l)`

List all tasks in the list.

Example:

```
tasker list
```

### `tasker done`

Mark a task as done. You will be prompted for the task ID.

Example:

```
tasker done <task_name>
```

### `tasker rm`

Delete a task from the list. You will be prompted for the task ID.

Example:

```
tasker rm <task_name>
```

### `tasker --help (-h)`

Show help information about the available commands.

Example:

```
tasker --help
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.