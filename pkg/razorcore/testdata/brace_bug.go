// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: cases/brace_bug.gohtml

package cases

import (
	"io"
	"strings"
)

// Brace_bug generates cases/brace_bug.gohtml
func Brace_bug() string {
	var _b strings.Builder
	RenderBrace_bug(&_b)
	return _b.String()
}

// RenderBrace_bug render cases/brace_bug.gohtml
func RenderBrace_bug(_buffer io.StringWriter) {

	isActive := func(name string) {
		if active == name {

			_buffer.WriteString("<li class=\"active\">\n        ")
		} else {

			_buffer.WriteString("<li>\n        ")
		}
	}

}
