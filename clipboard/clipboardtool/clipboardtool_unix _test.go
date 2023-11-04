//go:build freebsd || linux || netbsd || openbsd || solaris || dragonfly

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newClipboardTool(t *testing.T) {
	testCases := []struct {
		desc           string
		lookPathMock   func(file string) (string, error)
		expectedOutput *clipboardTool
		expectedError  error
	}{
		{
			desc: "xsel tool option is available",
			lookPathMock: func(toolName string) (string, error) {
				if toolName == xsel {
					return "/path/to/xsel", nil
				}
				return "", errors.New("not available")
			},
			expectedOutput: &clipboardTool{
				CopyTool: &CopyTool{
					Name:    xsel,
					CmdArgs: []string{"--input", "--clipboard"},
				},
				PasteTool: &PasteTool{
					Name:    xsel,
					CmdArgs: []string{"--output", "--clipboard"},
				},
			},
		},
		{
			desc: "xclip tool option is available",
			lookPathMock: func(toolName string) (string, error) {
				if toolName == xclip {
					return "/path/to/xclip", nil
				}
				return "", errors.New("not available")
			},
			expectedOutput: &clipboardTool{
				CopyTool: &CopyTool{
					Name:    xclip,
					CmdArgs: []string{"-in", "-selection", "clipboard"},
				},
				PasteTool: &PasteTool{
					Name:    xclip,
					CmdArgs: []string{"-out", "-selection", "clipboard"},
				},
			},
		},
		{
			desc: "wayland tools options are available",
			lookPathMock: func(toolName string) (string, error) {
				if toolName == wlcopy || toolName == wlpaste {
					return "", nil
				}
				return "", errors.New("not available")
			},
			expectedOutput: &clipboardTool{
				CopyTool: &CopyTool{
					Name: wlcopy,
				},
				PasteTool: &PasteTool{
					Name:    wlpaste,
					CmdArgs: []string{"--no-newline"},
				},
			},
		},
		{
			desc: "termux tools options are available",
			lookPathMock: func(toolName string) (string, error) {
				if toolName == termuxClipboardGet || toolName == termuxClipboardSet {
					return "", nil
				}
				return "", errors.New("not available")
			},
			expectedOutput: &clipboardTool{
				CopyTool: &CopyTool{
					Name: termuxClipboardSet,
				},
				PasteTool: &PasteTool{
					Name: termuxClipboardGet,
				},
			},
		},
		{
			desc: "no options available",
			lookPathMock: func(toolName string) (string, error) {
				return "", errors.New("not available")
			},
			expectedError: errors.New("no clipboard utilities available"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			lookPath = tc.lookPathMock
			ct, err := newClipboardTool()
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error to be %v, got nil", tc.expectedError)
				}
				require.NotNil(t, ct)
				require.Equal(t, tc.expectedOutput, ct)
			}
		})
	}
}
