//go:build darwin

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboard

import (
	"os/exec"

	"github.com/tiagomelo/go-clipboard/clipboard/clipboardtool"
	"github.com/tiagomelo/go-clipboard/clipboard/command"
)

// newCmd is a convenience function that creates a new command instance.
// It takes a command name and a variable number of arguments, then returns a command.Command
// which abstracts over the exec.Command for ease of testing and decoupling.
var newCmd = func(cmdName string, cmdArgs ...string) command.Command {
	return command.New(exec.Command(cmdName, cmdArgs...))
}

// copyText takes a string and copies it to the system clipboard.
// It uses the clipboardtool package to determine the appropriate tool and command package
// to execute the copy operation. An error is returned if the tool cannot be initialized or
// if the TextInput method fails.
func copyText(s string) error {
	ct, err := clipboardtool.New()
	if err != nil {
		return err
	}
	cmd := newCmd(ct.CopyTool.Name)
	return cmd.TextInput(s)
}

// pasteText retrieves text from the system clipboard.
// It uses the clipboardtool package to determine the appropriate tool and command package
// to execute the paste operation. It returns the pasted text and any error encountered.
func pasteText() (string, error) {
	ct, err := clipboardtool.New()
	if err != nil {
		return "", err
	}
	cmd := newCmd(ct.PasteTool.Name)
	return cmd.TextOutput()
}
