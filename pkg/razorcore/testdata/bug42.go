// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: cases/bug42.gohtml

package cases

import (
	"Tpl"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Bug42 generates cases/bug42.gohtml
func Bug42() string {
	var _b strings.Builder
	RenderBug42(&_b)
	return _b.String()
}

// RenderBug42 render cases/bug42.gohtml
func RenderBug42(_buffer io.StringWriter) {
	_buffer.WriteString("\n<div class=\"container\">\n    ")
	_buffer.WriteString(gorazor.HTMLEscape((Tpl.TplBread([]string{"选择邮寄方式"}))))
	_buffer.WriteString("\n</div>")

}
