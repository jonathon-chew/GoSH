# GoSH

`GoSH` is a small shell written in Go, with a growing set of Unix-style built-ins.

I built it to understand what a command-line shell actually has to do: read user input, recognise built-in commands, look up binaries on `PATH`, run external programs, and keep track of the current working directory.

This is a learning project rather than a replacement for an everyday shell.

## What It Does

- Starts an interactive prompt.
- Supports built-in commands including `cd`, `echo`, `ls`, `pwd`, and `type`.
- Searches the system `PATH` for executable binaries.
- Runs external commands with arguments.
- Keeps external command output separate from shell control flow.

## Built-In Commands

Current built-ins:

- `exit` closes the shell.
- `echo` prints text back to the terminal.
- `cd` changes the current working directory.
- `pwd` prints the current working directory.
- `type` reports whether a command is built in or resolves to an executable on `PATH`.

## Running It

From the project root:

```bash
go run ./cmd/GoSH
```

Build a binary:

```bash
go build ./cmd/GoSH
```

Then run it:

```bash
./GoSH
```

## Example Session

```text
$ pwd
/path/to/GoSH
$ echo hello
 hello
$ type go
go is in /usr/local/go/bin/go
$ exit
```

## Project Structure

```text
cmd/GoSH/          application entry point
    internal/builtins/  built-in command implementations
    internal/ls/        directory listing support
    internal/utils/     command lookup, execution, and history helpers
scripts/            helper scripts used across the project
```

## Current Status

The shell currently handles an interactive loop, command history, directory listing, several built-ins, and external command execution through Go's `os/exec` package. Input is currently split on spaces, so quoted arguments are not yet handled like a full POSIX shell.

The next useful improvements would be:

- move shell behavior out of `main.go` into testable packages,
- add command parsing that handles quoted arguments,
- add history support,
- add tests for command lookup and built-in behavior,
- improve errors around malformed commands such as `type` with no argument.

## What This Demonstrates

This project shows low-level CLI work in Go: terminal input, process execution, filesystem checks, environment-variable use, and the difference between shell built-ins and external binaries.

## License

See [LICENSE](./LICENSE) for details.
