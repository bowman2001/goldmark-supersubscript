package extension

import (
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

func TestSubscript(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			Subscript, Superscript,
			extension.Strikethrough,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/subscript.txt", t, testutil.ParseCliCaseArg()...)
}

func TestSubscriptDump(t *testing.T) {
	input := "The **formula H~2~O** describes the chemical composition of water."
	markdown := goldmark.New(
		goldmark.WithExtensions(
			Subscript, Superscript,
			extension.Strikethrough,
		),
	)
	root := markdown.Parser().Parse(text.NewReader([]byte(input)))
	root.Dump([]byte(input), 0)
	// Prints to stdout, so just test that it doesn't crash
}
