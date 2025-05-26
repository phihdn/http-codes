# http-codes

A simple Go application demonstrating how to use the [Bubble Tea](https://github.com/charmbracelet/bubbletea) terminal UI framework to create a web server status checker.

## Description

This application makes an HTTP request to charm.sh and displays the resulting HTTP status code. It showcases:

- Basic Bubble Tea application structure (Model-View-Update)
- HTTP client usage in Go
- Error handling
- Terminal user interface

## Usage

```go
go run main.go
```

The program will display the HTTP status of the target URL. Press Ctrl+C to exit.

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea): Terminal UI framework
- Go standard library (net/http, fmt, time)
