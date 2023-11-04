//go:build darwin

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboard

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/go-clipboard/clipboard/command"
)

func Test_copyText(t *testing.T) {
	testCases := []struct {
		desc          string
		mockClosure   func(m *mockCommand)
		expectedError error
	}{
		{
			desc: "happy path",
		},
		{
			desc: "error",
			mockClosure: func(m *mockCommand) {
				m.ErrTextInput = errors.New("text input error")
			},
			expectedError: errors.New("text input error"),
		},
	}
	for _, tc := range testCases {
		m := new(mockCommand)
		newCmd = func(cmdName string, cmdArgs ...string) command.Command {
			return m
		}
		t.Run(tc.desc, func(t *testing.T) {
			if tc.mockClosure != nil {
				tc.mockClosure(m)
			}
			err := copyText("some text")
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error to be %v, got nil", tc.expectedError)
				}
			}
		})
	}
}

func Test_pasteText(t *testing.T) {
	testCases := []struct {
		desc           string
		mockClosure    func(m *mockCommand)
		expectedOutput string
		expectedError  error
	}{
		{
			desc: "happy path",
			mockClosure: func(m *mockCommand) {
				m.Output = "some text"
			},
			expectedOutput: "some text",
		},
		{
			desc: "error",
			mockClosure: func(m *mockCommand) {
				m.ErrOutput = errors.New("paste text error")
			},
			expectedError: errors.New("paste text error"),
		},
	}
	for _, tc := range testCases {
		m := new(mockCommand)
		newCmd = func(cmdName string, cmdArgs ...string) command.Command {
			return m
		}
		t.Run(tc.desc, func(t *testing.T) {
			tc.mockClosure(m)
			output, err := pasteText()
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error to be %v, got nil", tc.expectedError)
				}
				require.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}

type mockCommand struct {
	ErrTextInput error
	ErrOutput    error
	Output       string
}

func (m *mockCommand) TextInput(text string) error {
	return m.ErrTextInput
}

func (m *mockCommand) TextOutput() (string, error) {
	return m.Output, m.ErrOutput
}
