package extension

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

func TestSuperscript(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Superscript,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/superscript.txt", t, testutil.ParseCliCaseArg()...)
}
