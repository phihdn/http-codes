// Package main provides a simple HTTP status checker application using Bubble Tea TUI framework
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// URL to check for HTTP status
const url = "https://charm.sh"

// model represents the application state
type model struct {
	status int   // HTTP status code received from the server
	err    error // Error that might occur during the HTTP request
}

// checkServer performs an HTTP GET request to the specified URL
// and returns either a statusMsg with the HTTP status code or an errMsg if an error occurs
func checkServer() tea.Msg {
	// Create an HTTP client with a 10-second timeout
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}
	defer res.Body.Close() // nolint:errcheck

	return statusMsg(res.StatusCode)
}

// Custom message types for the Bubble Tea framework
type (
	statusMsg int                 // Carries the HTTP status code
	errMsg    struct{ err error } // Wraps an error message
)

// Error implements the error interface for errMsg
func (e errMsg) Error() string {
	return e.err.Error()
}

// Init is the first function that will be called when the program starts
// It returns the checkServer command as the first action to perform
func (m model) Init() tea.Cmd {
	return checkServer
}

// Update handles incoming messages and updates the model accordingly
// It implements the Bubble Tea update loop
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		// When we receive a status code, store it and quit the application
		m.status = int(msg)
		return m, tea.Quit
	case errMsg:
		// When we receive an error, store it and quit the application
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		/*
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
			equals...
		*/
		// Allow the user to quit the application with Ctrl+C
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	// Return the model unchanged if no specific message was handled
	return m, nil
}

// View renders the current state of the application to a string
// This string will be displayed in the terminal
func (m model) View() string {
	if m.err != nil {
		// Display error message if an error occurred
		return fmt.Sprintf("Error: %v\nPress Ctrl+C to exit.", m.err)
	}

	// Start building the status message
	s := fmt.Sprintf("Checking %s...\n", url)

	if m.status > 0 {
		// If we have a status code, display it with its text representation
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}

	// Add some spacing for better readability
	return "\n" + s + "\n\n"
}

// main is the entry point of the application
// It creates and runs a new Bubble Tea program with our model
func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
