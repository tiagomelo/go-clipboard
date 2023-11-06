package clipboard_test

import (
	"fmt"
	"os"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

func Example() {
	text := "some text"
	c := clipboard.New()
	if err := c.CopyText(text); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("text \"%s\" was copied into clipboard. Paste it elsewhere.\n", text)

	pastedText, err := c.PasteText()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("text from clipboard: %v\n", pastedText)
}
