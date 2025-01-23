// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

func main() {
	text := "some text"
	c := clipboard.New()

	if len(os.Args) > 0 {
		c = clipboard.New(clipboard.ClipboardOptions{Primary: true})
	}

	if err := c.CopyText(text); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("text \"%s\" was copied into clipboard. Paste it elsewhere.\n", text)
}
