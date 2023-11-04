//go:build windows

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package command

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_textInput(t *testing.T) {
	testCases := []struct {
		desc          string
		mockClosure   func(c *mockSysCmd, w *mockIoPipeWriter)
		expectedError error
	}{
		{
			desc: "happy path",
		},
		{
			desc: "error when getting pipe for command",
			mockClosure: func(c *mockSysCmd, w *mockIoPipeWriter) {
				c.ErrStdinPipe = errors.New("pipe error")
			},
			expectedError: errors.New("getting pipe for command: pipe error"),
		},
		{
			desc: "error when starting command",
			mockClosure: func(c *mockSysCmd, w *mockIoPipeWriter) {
				c.ErrStart = errors.New("start error")
			},
			expectedError: errors.New("starting command: start error"),
		},
		{
			desc: "error when writing input for command",
			mockClosure: func(c *mockSysCmd, w *mockIoPipeWriter) {
				w.WriteErr = errors.New("write error")
			},
			expectedError: errors.New("writing input for command: write error"),
		},
		{
			desc: "error when closing input for command",
			mockClosure: func(c *mockSysCmd, w *mockIoPipeWriter) {
				w.CloseErr = errors.New("close error")
			},
			expectedError: errors.New("closing input: close error"),
		},
		{
			desc: "error when waiting for command",
			mockClosure: func(c *mockSysCmd, w *mockIoPipeWriter) {
				c.ErrWait = errors.New("wait error")
			},
			expectedError: errors.New("waiting for command: wait error"),
		},
	}
	for _, tc := range testCases {
		c := new(mockSysCmd)
		w := new(mockIoPipeWriter)
		c.IoPipeWriterMock = w
		t.Run(tc.desc, func(t *testing.T) {
			if tc.mockClosure != nil {
				tc.mockClosure(c, w)
			}
			err := textInput(c, "some text")
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

func Test_textOutput(t *testing.T) {
	testCases := []struct {
		desc           string
		mockClosure    func(c *mockSysCmd)
		expectedOutput string
		expectedError  error
	}{
		{
			desc: "happy path",
			mockClosure: func(c *mockSysCmd) {
				c.CmdOutput = []byte("some output")
			},
			expectedOutput: "some output",
		},
		{
			desc: "error when getting command output",
			mockClosure: func(c *mockSysCmd) {
				c.ErrOutput = errors.New("output error")
			},
			expectedError: errors.New("getting output for command: output error"),
		},
	}
	for _, tc := range testCases {
		c := new(mockSysCmd)
		t.Run(tc.desc, func(t *testing.T) {
			tc.mockClosure(c)
			output, err := textOutput(c)
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

type mockIoPipeWriter struct {
	N        int
	WriteErr error
	CloseErr error
}

func (i *mockIoPipeWriter) Write(p []byte) (n int, err error) {
	return i.N, i.WriteErr
}

func (i *mockIoPipeWriter) Close() error {
	return i.CloseErr
}

type mockSysCmd struct {
	ErrStart         error
	ErrOutput        error
	ErrStdinPipe     error
	ErrWait          error
	CmdOutput        []byte
	IoPipeWriterMock *mockIoPipeWriter
}

func (m *mockSysCmd) Start() error {
	return m.ErrStart
}

func (m *mockSysCmd) Output() ([]byte, error) {
	return m.CmdOutput, m.ErrOutput
}

func (m *mockSysCmd) StdinPipe() (ioPipeWriter, error) {
	return m.IoPipeWriterMock, m.ErrStdinPipe
}

func (m *mockSysCmd) Wait() error {
	return m.ErrWait
}
