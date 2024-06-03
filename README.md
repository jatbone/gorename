# gorename

`gorename` is a command-line tool written in Go for batch renaming files based on a regex pattern and replacement string. It also supports undoing renaming operations by maintaining a log of changes.

## Features

- Rename files using a regex pattern.
- Dry run mode to preview changes without applying them.
- Recursive renaming through subdirectories.
- Undo renaming operations using a log file.

## Installation

To install `gorename`, ensure you have Go installed and then run:

```sh
go get github.com/jatbone/gorename
```

## Usage

The main entry point for the application is in `cmd/gorename/main.go`. The tool provides several command-line flags for its operations.

### Command-line Flags

- `-pattern`: (Required) Regex pattern to match filenames.
- `-replacement`: (Required) Replacement format for filenames.
- `-path`: (Optional) Path to the files. Defaults to the current directory.
- `-dryRun`: (Optional) Perform a dry run without renaming files.
- `-recursive`: (Optional) Recursively rename files in subdirectories.
- `-undo`: (Optional) Undo rename operations.
- `-undoLogFile`: (Optional) Specify a custom log file for undo operations.

### Examples

#### Basic Rename

```sh
gorename -pattern "testfile(.*)" -replacement "replace$1" -path "./files"
```

This command will rename all files matching `testfile(.*)` to `replace$1` in the `./files` directory.

#### Dry Run

```sh
gorename -pattern "testfile(.*)" -replacement "replace$1" -path "./files" -dryRun
```

This command will display the changes without actually renaming the files.

#### Recursive Rename

```sh
gorename -pattern "testfile(.*)" -replacement "replace$1" -path "./files" -recursive
```

This command will rename files matching the pattern in `./files` and all its subdirectories.

#### Undo Rename

```sh
gorename -undo
```

This command will undo the last renaming operation using the default log file `gorename-log.txt`.

#### Custom Log File

```sh
gorename -undo -undoLogFile "./custom-log.txt"
```

This command will undo the renaming operations recorded in `./custom-log.txt`.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
