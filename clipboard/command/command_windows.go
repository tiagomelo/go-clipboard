//go:build windows

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package command

import "github.com/pkg/errors"

// textInput sends the provided text as input to the system command.
// It wraps the underlying system call processes with additional error handling.
// It returns an error if any step of the command execution process fails.
func textInput(c sysCommand, text string) error {
	in, err := c.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "getting pipe for command")
	}
	if err := c.Start(); err != nil {
		return errors.Wrap(err, "starting command")
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return errors.Wrap(err, "writing input for command")
	}
	if err := in.Close(); err != nil {
		return errors.Wrap(err, "closing input")
	}
	if err := c.Wait(); err != nil {
		return errors.Wrap(err, "waiting for command")
	}
	return nil
}

// textOutput executes the system command and captures its standard output.
// It returns the captured output as a string along with any error that occurred during execution.
func textOutput(c sysCommand) (string, error) {
	out, err := c.Output()
	if err != nil {
		return "", errors.Wrap(err, "getting output for command")
	}
	return string(out), nil
}
