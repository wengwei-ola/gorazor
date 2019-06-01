package cases

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

func Bug9(l *Locale) string {
	var _b strings.Builder
	RenderBug9(&_b, l)
	return _b.String()
}

func RenderBug9(_buffer io.StringWriter, l *Locale) {
	_buffer.WriteString("\n<span>")
	_buffer.WriteString(gorazor.HTMLEscape(l.T(`for`)))
	_buffer.WriteString("</span>")

}
