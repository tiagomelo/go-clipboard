//go:build darwin

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import (
	"errors"
	"os/exec"
)

const (
	// pbcopy is the name of the macOS command-line tool used to copy text to the clipboard.
	pbcopy = "pbcopy"
	// pbpaste is the name of the macOS command-line tool used to paste text from the clipboard.
	pbpaste = "pbpaste"
)

var (
	// copyTool is a preconfigured CopyTool for macOS using the pbcopy utility.
	copyTool = &CopyTool{
		Name: pbcopy,
	}
	// pasteTool is a preconfigured PasteTool for macOS using the pbpaste utility.
	pasteTool = &PasteTool{
		Name: pbpaste,
	}
	// lookPath is a variable holding the exec.LookPath function,
	// used to check for the presence of a command in the system's PATH.
	lookPath = exec.LookPath

	errNoCopyUtilitiesFound  = errors.New("no clipboard copy utilities available")
	errNoPasteUtilitiesFound = errors.New("no clipboard paste utilities available")
)

// newClipboardTool initializes a new clipboardTool instance by
// checking the availability of clipboard utilities.
func newClipboardTool(primary bool) (*clipboardTool, error) {
	if isAvailable := isToolAvailable(copyTool.Name); !isAvailable {
		return nil, errNoCopyUtilitiesFound
	}
	if isAvailable := isToolAvailable(pasteTool.Name); !isAvailable {
		return nil, errNoPasteUtilitiesFound
	}
	return &clipboardTool{
		CopyTool:  copyTool,
		PasteTool: pasteTool,
	}, nil
}

// isToolAvailable checks if a clipboard utility tool
// is available in the system's PATH.
func isToolAvailable(toolName string) bool {
	if _, err := lookPath(toolName); err != nil {
		return false
	}
	return true
}
