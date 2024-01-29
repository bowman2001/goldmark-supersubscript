package extension

import (
	"github.com/yuin/goldmark/extension"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

func TestSubscript(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Subscript,
			extension.Strikethrough,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/subscript.txt", t, testutil.ParseCliCaseArg()...)
}
