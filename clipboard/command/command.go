// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package command

import (
	"io"
	"os/exec"
)

// ioPipeWriter is an interface that abstracts the io.WriteCloser interface
// to allow for mocking the standard input pipe of a system command.
type ioPipeWriter interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ioPipeWriterWrapper wraps an io.WriteCloser to conform to the ioPipeWriter interface.
type ioPipeWriterWrapper struct {
	wc io.WriteCloser
}

// Write writes data to the underlying io.WriteCloser.
func (i *ioPipeWriterWrapper) Write(p []byte) (n int, err error) {
	return i.wc.Write(p)
}

// Close closes the underlying io.WriteCloser.
func (i *ioPipeWriterWrapper) Close() error {
	return i.wc.Close()
}

// sysCommand is an interface that abstracts the methods of exec.Cmd
// that are used within the command package, allowing for mocking in tests.
type sysCommand interface {
	Start() error
	Output() ([]byte, error)
	StdinPipe() (ioPipeWriter, error)
	Wait() error
}

// sysCommandWrapper wraps an exec.Cmd to conform to the sysCommand interface.
type sysCommandWrapper struct {
	cmd *exec.Cmd
}

// Start starts the specified command but does not wait for it to complete.
func (sc *sysCommandWrapper) Start() error {
	return sc.cmd.Start()
}

// Output runs the command and returns its standard output.
func (sc *sysCommandWrapper) Output() ([]byte, error) {
	return sc.cmd.Output()
}

// StdinPipe returns a pipe that will be connected to the command's standard input
// when the command starts.
func (sc *sysCommandWrapper) StdinPipe() (ioPipeWriter, error) {
	p, err := sc.cmd.StdinPipe()
	return &ioPipeWriterWrapper{p}, err
}

// Wait waits for the command to exit and waits for any copying to stdin or
// copying from stdout or stderr to complete.
func (sc *sysCommandWrapper) Wait() error {
	return sc.cmd.Wait()
}

// Command is an interface that provides methods for sending text input to a command
// and receiving text output from a command.
type Command interface {
	TextInput(text string) error
	TextOutput() (string, error)
}

// command is an implementation of the Command interface that uses sysCommand
// to execute system commands and process their input and output.
type command struct {
	sc sysCommand
}

// New creates a new command instance with the specified exec.Cmd.
func New(cmd *exec.Cmd) Command {
	return &command{
		sc: &sysCommandWrapper{cmd},
	}
}

// TextInput sends the provided text as input to the system command.
func (c *command) TextInput(text string) error {
	return textInput(c.sc, text)
}

// TextOutput executes the command and returns its output as a string.
func (c *command) TextOutput() (string, error) {
	return textOutput(c.sc)
}
