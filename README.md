# go-sh

`go-sh` is a small shell written in Go.

I built it to understand what a command-line shell actually has to do: read user input, recognise built-in commands, look up binaries on `PATH`, run external programs, and keep track of the current working directory.

This is a learning project rather than a replacement for an everyday shell.

## What It Does

- Starts an interactive prompt.
- Supports a small set of built-in commands.
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
go run ./cmd/go-sh
```

Build a binary:

```bash
go build ./cmd/go-sh
```

Then run it:

```bash
./go-sh
```

## Example Session

```text
$ pwd
/path/to/go-sh
$ echo hello
 hello
$ type go
go is in /usr/local/go/bin/go
$ exit
```

## Project Structure

```text
cmd/go-sh/          application entry point
internal/cli/       reserved for future CLI/shared shell logic
scripts/            helper scripts used across the project
```

## Current Status

The shell currently handles a small interactive loop and a few core built-ins. External command execution works through Go's `os/exec` package.

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
