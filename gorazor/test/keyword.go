package cases

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

func Keyword() string {
	var _b strings.Builder
	RenderKeyword(&_b)
	return _b.String()
}

func RenderKeyword(_buffer io.StringWriter) {
	_buffer.WriteString("BLK(<span>rememberingsteve@apple.com ")
	_buffer.WriteString(gorazor.HTMLEscape(username))
	_buffer.WriteString("</span>)BLK")

}
