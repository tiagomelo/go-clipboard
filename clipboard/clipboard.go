// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboard

// clipboard is an unexported type that implements the Clipboard interface.
type clipboard struct {
	Primary bool
}

// internal clipboard target flag
var usePrimary bool

// exported flag container
type ClipboardOptions struct {
	Primary bool
}

// Clipboard is the interface that wraps the basic clipboard operations.
type Clipboard interface {
	// CopyText takes a string and copies it to the system clipboard.
	// It returns an error if the copying process fails.
	CopyText(s string) error

	// PasteText retrieves text from the system clipboard.
	// It returns the text as a string and an error if the paste operation fails.
	PasteText() (string, error)
}

// New creates and returns a new Clipboard instance that can be used
// to interact with the system clipboard.
func New(opts ...ClipboardOptions) Clipboard {
	cb := &clipboard{}

	if len(opts) == 1 {
		usePrimary = opts[0].Primary
	}

	return cb
}

// CopyText implements the Clipboard interface's CopyText method.
// It calls the copyText function to perform the actual operation.
func (c *clipboard) CopyText(s string) error {
	return copyText(s)
}

// PasteText implements the Clipboard interface's PasteText method.
// It calls the pasteText function to perform the actual operation.
func (c *clipboard) PasteText() (string, error) {
	return pasteText()
}
