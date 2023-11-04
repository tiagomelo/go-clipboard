//go:build freebsd || linux || netbsd || openbsd || solaris || dragonfly

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import (
	"errors"
	"os/exec"
)

const (
	// xsel is a clipboard utility for X11.
	xsel = "xsel"
	// xclip is another clipboard utility for X11.
	xclip = "xclip"

	// wlcopy is a clipboard utility for the Wayland display server.
	wlcopy = "wl-copy"
	// wlpaste is a clipboard utility for the Wayland display server.
	wlpaste = "wl-paste"

	// termuxClipboardGet is a clipboard utility for Termux, an Android terminal emulator.
	termuxClipboardGet = "termux-clipboard-get"
	// termuxClipboardSet is a clipboard utility for Termux, an Android terminal emulator.
	termuxClipboardSet = "termux-clipboard-set"
)

var (
	// copyTools is a list of available CopyTool configurations for different environments.
	copyTools = []*CopyTool{
		{
			Name:    xsel,
			CmdArgs: []string{"--input", "--clipboard"},
		},
		{
			Name:    xclip,
			CmdArgs: []string{"-in", "-selection", "clipboard"},
		},
		{
			Name: wlcopy,
		},
		{
			Name: termuxClipboardSet,
		},
	}
	// pasteTools is a list of available PasteTool configurations for different environments.
	pasteTools = []*PasteTool{
		{
			Name:    xsel,
			CmdArgs: []string{"--output", "--clipboard"},
		},
		{
			Name:    xclip,
			CmdArgs: []string{"-out", "-selection", "clipboard"},
		},
		{
			Name:    wlpaste,
			CmdArgs: []string{"--no-newline"},
		},
		{
			Name: termuxClipboardGet,
		},
	}
	// lookPath is a variable holding the exec.LookPath function,
	// used to check for the presence of a command in the system's PATH.
	lookPath = exec.LookPath

	errNoUtilitiesFound = errors.New("no clipboard utilities available")
)

// newClipboardTool selects the first available pair of
// copy and paste tools from the predefined list.
func newClipboardTool() (*clipboardTool, error) {
	for i, ct := range copyTools {
		pt := pasteTools[i]
		if available := toolsAreAvailable(ct.Name, pt.Name); available {
			return &clipboardTool{
				CopyTool:  ct,
				PasteTool: pt,
			}, nil
		}
	}
	return nil, errNoUtilitiesFound
}

// toolsAreAvailable checks for the existence of the specified
// tools by name in the system's PATH.
func toolsAreAvailable(toolNames ...string) bool {
	for _, toolName := range toolNames {
		if _, err := lookPath(toolName); err != nil {
			return false
		}
	}
	return true
}
