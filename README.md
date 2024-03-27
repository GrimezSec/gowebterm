# GoWebTerm

GoWebTerm is a simple web application developed using the Go programming language. This application provides a terminal-like interface accessible through a web browser, allowing users to input and execute specific commands.

## Features

- Terminal-like interface accessible through a web browser.
- Support for pressing Enter key for command input.
- Real-time output updates.
- User-friendly interface.

## How to Use

1. Go compiler is required to run the application.
2. In the root directory of the application, run the command `go run main.go`.
3. Open your web browser and navigate to `http://localhost:8080` (Default port is 8080).
4. Enter the desired command in the command input box and press Enter or click the "Execute" button.

## Requirements

- Go compiler is required.
- Gorilla WebSocket library is needed. To install:
  ` go get github.com/gorilla/websocket `
