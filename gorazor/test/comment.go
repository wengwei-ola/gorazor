package cases

import (
	"bytes"
	"io"
	"strings"
)

func Comment() string {
	var _b strings.Builder
	RenderComment(&_b)
	return _b.String()
}

func RenderComment(_buffer io.StringWriter) {
	_buffer.WriteString("\n\n\n\n<p>hello </p>")

	hello

}
