# Go WinService Template

This is a template to create a Windows Service in Go.

## How to use

Main loop occurs in `/lib/app/app.go`

```go
log.Info("Running")
time.Sleep(5 * time.Second)
```

For now, it just logs a message and waits 5 seconds.

## Development

Use `go run main.go` to run the service quickly.

Use `build` to build the following files to the `/bin`:

- **main.exe**
- install.bat
- uninstall.bat
- start.bat
- stop.bat

## Install

Install needs the target install path and the service name.

```cmd
install "C:\Program Files\MyService" "MyService"
```

## Uninstall

Uninstall simply removes the service and the files at the install path.

```cmd
uninstall
```

## Start

Starts the service.

```cmd
start
```

## Stop

Stops the service.

```cmd
stop
```
