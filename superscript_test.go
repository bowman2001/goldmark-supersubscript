package extension

import (
	"github.com/yuin/goldmark/text"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

func TestSuperscript(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			Superscript, Subscript,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/superscript.txt", t, testutil.ParseCliCaseArg()...)
}

func TestSuperscriptDump(t *testing.T) {
	input := "Parabola: f(x) = x^2^. Amazing"
	markdown := goldmark.New(
		goldmark.WithExtensions(
			Superscript, Subscript,
		),
	)
	root := markdown.Parser().Parse(text.NewReader([]byte(input)))
	root.Dump([]byte(input), 0)
	// Prints to stdout, so just test that it doesn't crash
}
