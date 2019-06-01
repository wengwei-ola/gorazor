package cases

import (
	"bytes"
	"cases/layout"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"kp/models"
	"strings"
)

func Edit(u *models.User) string {
	var _b strings.Builder
	WriteEdit(&_b, u)
	return _b.String()
}

func WriteEdit(_buffer io.StringWriter, u *models.User) {

	_body := func(_buffer io.StringWriter) {
		_buffer.WriteString("\n<div style=\"width: 500px\">\n<form role=\"form\">\n  <div class=\"form-group\">\n    <label for=\"exampleInputEmail1\">名字</label>\n    <input type=\"email\" class=\"form-control\" id=\"exampleInputEmail1\" placeholder=\"Enter email\" value=\"")
		_buffer.WriteString(gorazor.HTMLEscape(u.Name))
		_buffer.WriteString("\">\n  </div>\n  <div class=\"form-group\">\n    <label for=\"exampleInputPassword1\">电邮</label>\n    <input type=\"email\" class=\"form-control\" id=\"exampleInputPassword1\" placeholder=\"电邮\" value=\"")
		_buffer.WriteString(gorazor.HTMLEscape(u.Email))
		_buffer.WriteString("\">\n  </div>\n  <button type=\"submit\" class=\"btn btn-primary\">保存</button>\n  <a href=\"/admin/user\" class=\"btn btn-default pull-right\">返回</a>\n</form>\n</div>")

	}

	title := func(_buffer io.StringWriter) {

		_buffer.WriteString("用户管理")

	}

	return layout.Base(_buffer, body, title, nil)
}
