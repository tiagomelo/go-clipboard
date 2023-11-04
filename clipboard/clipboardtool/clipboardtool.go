// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

// CopyTool encapsulates the details of a clipboard copy command.
type CopyTool struct {
	Name    string   // Name of the copy command or executable
	CmdArgs []string // Arguments required for the copy operation
}

// PasteTool encapsulates the details of a clipboard paste command.
type PasteTool struct {
	Name    string   // Name of the paste command or executable
	CmdArgs []string // Arguments required for the paste operation
}

// clipboardTool combines CopyTool and PasteTool to provide a unified interface
// for clipboard operations. It abstracts the underlying command-line tools used
// to interact with the system clipboard.
type clipboardTool struct {
	CopyTool  *CopyTool  // Tool to copy content to the clipboard
	PasteTool *PasteTool // Tool to paste content from the clipboard
}

// New initializes and returns a new instance of clipboardTool.
// It determines the appropriate tools to use based on the current system environment
// and returns an error if no suitable tools are found.
func New() (*clipboardTool, error) {
	return newClipboardTool()
}
