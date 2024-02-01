package extension

import (
	"github.com/yuin/goldmark/extension"
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
