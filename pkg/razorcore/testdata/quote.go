// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: cases/quote.gohtml

package cases

import (
	"io"
	"strings"
)

// Quote generates cases/quote.gohtml
func Quote() string {
	var _b strings.Builder
	RenderQuote(&_b)
	return _b.String()
}

// RenderQuote render cases/quote.gohtml
func RenderQuote(_buffer io.StringWriter) {
	_buffer.WriteString("<html>'text'</html>")

}
