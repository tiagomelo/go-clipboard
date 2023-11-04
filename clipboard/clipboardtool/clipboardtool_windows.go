//go:build windows

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import (
	"errors"
	"os/exec"
)

const (
	// clip is the name of the Windows command-line tool used to copy text to the clipboard.
	clip = "clip.exe"
	// powershell is used to access advanced Windows functionality,
	// including the ability to paste text from the clipboard.
	powershell = "powershell"
)

var (
	// copyTool is a preconfigured CopyTool for Windows using the clip utility.
	copyTool = &CopyTool{
		Name: clip,
	}
	// pasteTool is a preconfigured PasteTool for Windows using PowerShell commands.
	pasteTool = &PasteTool{
		Name:    powershell,
		CmdArgs: []string{"Get-Clipboard"},
	}
	// lookPath is a variable holding the exec.LookPath function,
	// used to check for the presence of a command in the system's PATH.
	lookPath = exec.LookPath

	errNoCopyUtilitiesFound  = errors.New("no clipboard copy utilities available")
	errNoPasteUtilitiesFound = errors.New("no clipboard paste utilities available")
)

// newClipboardTool checks the availability of clipboard utilities
// and initializes a new clipboardTool.
func newClipboardTool() (*clipboardTool, error) {
	if isAvailable := toolIsAvailable(copyTool.Name); !isAvailable {
		return nil, errNoCopyUtilitiesFound
	}
	if isAvailable := toolIsAvailable(pasteTool.Name); !isAvailable {
		return nil, errNoPasteUtilitiesFound
	}
	return &clipboardTool{
		CopyTool:  copyTool,
		PasteTool: pasteTool,
	}, nil
}

// toolIsAvailable verifies the presence of a clipboard utility in the system's PATH.
func toolIsAvailable(toolName string) bool {
	if _, err := lookPath(toolName); err != nil {
		return false
	}
	return true
}
