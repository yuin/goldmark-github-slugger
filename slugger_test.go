package slugger_test

import (
	"bytes"
	"testing"

	slugger "github.com/yuin/goldmark-github-slugger"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
	"github.com/yuin/goldmark/v2/util"
)

func TestSlugger(t *testing.T) {
	// taken from <https://github.com/gohugoio/hugo/blob/d6e6f9e500eebdeae8e28de830fff2e3bfc7d534/markup/goldmark/autoid_test.go>
	tests := []struct {
		title string
		slug  string
	}{
		{
			title: "God is good: 神真美好",
			slug:  "god-is-good-神真美好",
		},
		{
			title: "aßa",
			slug:  "aßa",
		},
		{
			title: "Number 32",
			slug:  "number-32",
		},
		{
			title: "Question?",
			slug:  "question",
		},
		{
			title: "1+2=3",
			slug:  "123",
		},
		{
			title: "Special !\"#$%&(parens)=?´* chars",
			slug:  "special-parens-chars",
		},
		{
			title: "Resumé",
			slug:  "resumé",
		},
		{
			title: "One-Hyphen",
			slug:  "one-hyphen",
		},
		{
			title: "Multiple--Hyphens",
			slug:  "multiple--hyphens",
		},
		{
			title: "Trailing hyphen-",
			slug:  "trailing-hyphen-",
		},
		{
			title: "Many   spaces  here",
			slug:  "many---spaces--here",
		},
		{
			title: "Forward/slash",
			slug:  "forwardslash",
		},
		{
			title: "Backward\\slash",
			slug:  "backwardslash",
		},
		{
			title: "Under_score",
			slug:  "under_score",
		},
		{
			title: "Nonbreaking Space", // Nonbreaking space U+00A0
			slug:  "nonbreakingspace",
		},
	}

	p := parser.New(parser.WithExtensions(slugger.Parser), parser.WithAutoHeadingID())
	r := html.New()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			var buf bytes.Buffer
			source := []byte("# " + tt.title)
			node := p.Parse(source)
			_ = r.Render(&buf, source, node)
			expected := `<h1 id="` + tt.slug + `">` + string(util.EscapeHTML([]byte(tt.title))) + `</h1>` + "\n"
			if buf.String() != expected {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}
