//go:build windows

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
			desc: "both tools are available",
			lookPathMock: func(toolName string) (string, error) {
				return "", nil
			},
			expectedOutput: &clipboardTool{
				CopyTool: &CopyTool{
					Name: clip,
				},
				PasteTool: &PasteTool{
					Name:    powershell,
					CmdArgs: []string{"Get-Clipboard"},
				},
			},
		},
		{
			desc: "clip is not available",
			lookPathMock: func(toolName string) (string, error) {
				if toolName == clip {
					return "", errors.New("not found")
				}
				return "", nil
			},
			expectedError: errors.New("no clipboard copy utilities available"),
		},
		{
			desc: "powershell is not available",
			lookPathMock: func(toolName string) (string, error) {
				if toolName == powershell {
					return "", errors.New("not found")
				}
				return "", nil
			},
			expectedError: errors.New("no clipboard paste utilities available"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			lookPath = tc.lookPathMock
			ct, err := newClipboardTool(false)
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
