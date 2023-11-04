// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"

	clipboard "github.com/tiagomelo/go-clipboard/clipboard"
)

func main() {
	c := clipboard.New()
	text, err := c.PasteText()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("text from clipboard: %v\n", text)
}
