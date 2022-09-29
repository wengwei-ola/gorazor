# gorazor

[![Build Status](https://travis-ci.org/sipin/gorazor.svg?branch=master)](https://travis-ci.org/sipin/gorazor)
[![Go Report Card](https://goreportcard.com/badge/github.com/sipin/gorazor?v=1)](https://goreportcard.com/report/github.com/sipin/gorazor)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-90%25-brightgreen.svg?longCache=true&style=flat)</a>
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

`gorazor` is the Go port of the razor view engine originated from [asp.net in 2011](http://weblogs.asp.net/scottgu/archive/2010/07/02/introducing-razor.aspx). In summary, `gorazor` is:

* Extremely Fast. Templates are converted into Go code and then compiled with optimizations.
* Concise syntax, no delimiter like `<?`, `<%`, or `{{`.
  * Original [Razor Syntax](http://www.asp.net/web-pages/tutorials/basics/2-introduction-to-asp-net-web-programming-using-the-razor-syntax) & [quick reference](http://haacked.com/archive/2011/01/06/razor-syntax-quick-reference.aspx/) for asp.net.
* Able to mix go code in view template
  * Insert code block to import & call arbitrary go modules & functions
  * Flow controls are just Go, no need to learn another mini-language
* Code generation approach
  * No reflection overhead
  * Go compiler validation for free
* Strong type view model
* Embedding templates support
* Layout/Section support

# Extremely Fast

`gorazor` is about **20X** times faster than [html/template](https://golang.org/pkg/html/template/) when using standard `strings.Builder` for template writing.

When using `quicktemplate`'s `ByteBuffer` and `unsafeStrToBytes` method to for template writing, `gorazor`'s performance is comparable to [quicktemplate](https://github.com/valyala/quicktemplate), if not faster.

Benchmark results:
```bash
$ go test -bench='Benchmark(Razor|RazorQuick|Quick|HTML)Template' -benchmem github.com/valyala/quicktemplate/tests github.com/sipin/gorazor/tests
goos: windows
goarch: amd64
pkg: github.com/valyala/quicktemplate/tests
BenchmarkQuickTemplate1-8       50000000                35.6 ns/op             0 B/op          0 allocs/op
BenchmarkQuickTemplate10-8      10000000               152 ns/op               0 B/op          0 allocs/op
BenchmarkQuickTemplate100-8      1000000              1460 ns/op               0 B/op          0 allocs/op
BenchmarkHTMLTemplate1-8         2000000               712 ns/op             608 B/op         21 allocs/op
BenchmarkHTMLTemplate10-8         500000              3586 ns/op            2834 B/op        111 allocs/op
BenchmarkHTMLTemplate100-8         50000             35180 ns/op           28055 B/op       1146 allocs/op
PASS
ok      github.com/valyala/quicktemplate/tests  11.360s
goos: windows
goarch: amd64
pkg: github.com/sipin/gorazor/tests
BenchmarkRazorTemplate1-8                       30000000                49.8 ns/op           224 B/op          3 allocs/op
BenchmarkRazorTemplate10-8                      10000000               122 ns/op             480 B/op          4 allocs/op
BenchmarkRazorTemplate100-8                      2000000               931 ns/op            4064 B/op          7 allocs/op
BenchmarkRazorQuickTemplate1-8                  100000000               19.9 ns/op             0 B/op          0 allocs/op
BenchmarkRazorQuickTemplate10-8                 20000000                82.5 ns/op             0 B/op          0 allocs/op
BenchmarkRazorQuickTemplate100-8                 2000000               767 ns/op               0 B/op          0 allocs/op
BenchmarkRazorQuickTemplateOriginal1-8          100000000               17.4 ns/op             0 B/op          0 allocs/op
BenchmarkRazorQuickTemplateOriginal10-8         20000000                68.8 ns/op             0 B/op          0 allocs/op
BenchmarkRazorQuickTemplateOriginal100-8         2000000               656 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/sipin/gorazor/tests  19.921s
```

* `BenchmarkRazorQuickTemplate`'s manually modified ensure **exact output** as quicktemplate for comparism.
* `BenchmarkRazorQuickTemplateOriginal` are `gorazor`'s default code-gen, which produce less white-space, thus faster.

# Usage

`gorazor` supports `go 1.10` and above, for go version **below 1.10**, you may use [gorazor classic version](https://github.com/sipin/gorazor/releases/tag/v1.0).

`go 1.12` are recommended for better **compiler optimization**.

## Install

```sh
go install github.com/sipin/gorazor@latest
```

## Usage

* Process folder: `gorazor template_folder output_folder`
* Process file: `gorazor template_file output_file`

## Examples

[Examples] gives examples using layout / helper etc.

When using layout, you may need to set `-prefix` parameter, like:

```bash
git clone https://github.com/sipin/gorazor/
cd gorazor
go build

# -prefix parameter here tells gorazor the current folder is the base path for github.com/sipin/gorazor
# So that, when importing "github.com/sipin/gorazor/examples/tpl/layout" in example/tpl/home.gohtml
# gorazor will know how to find the layout/helper files
./gorazor -prefix github.com/sipin/gorazor ./examples/tpl ./examples/tpl
go run example/main.go
```

# Syntax

## Variable

* `@variable` to insert **string** variable into html template
  * variable could be wrapped by arbitrary go functions
  * variable inserted will be automatically [escaped](http://golang.org/pkg/html/template/#HTMLEscapeString)

```html
<div>Hello @user.Name</div>
```

```html
<div>Hello @strings.ToUpper(req.CurrentUser.Name)</div>
```

Use `raw` to skip escaping:

```html
<div>@raw(user.Name)</div>
```

Only use `raw` when you are 100% sure what you are doing, please always be aware of [XSS attack](http://en.wikipedia.org/wiki/Cross-site_scripting).

## Flow Control

```php
@if .... {
	....
}

@if .... {
	....
} else {
	....
}

@for .... {

}

@{
	switch .... {
	case ....:
	      <p>...</p>
	case 2:
	      <p>...</p>
	default:
	      <p>...</p>
	}
}
```

Please use [example](https://github.com/sipin/gorazor/blob/master/examples/tpl/home.gohtml) for reference.

## Code block

Arbitrary go code could be used in templates, like creating a new variable.

```html
@{
	username := u.Name
	if u.Email != "" {
		username += "(" + u.Email + ")"
	}
}
<div class="welcome">
<h4>Hello @username</h4>
</div>
```

It's recommended to keep clean separation of code & view. **Please consider** move logic into your code before creating a code block in a template.

## Declaration

The **first code block** in template is strictly for declaration:

* imports
* model type
* layout

like:

```go
@{
	import  (
		"kp/models"   //import `"kp/models"` package
		"tpl/layout"  //import tpl/layout namespace
	)

	layout := layout.Base //Use layout package's **Base func** for layout
	var user *models.User //1st template param
	var blog *models.Blog //2nd template param
}
...
```

**first code block** must be at the beginning of the template, i.e. before any html.

Any other codes inside the first code block will **be ignored**.

import must be wrapped in `()`, `import "package_name"` is not yet supported.

The variables declared in **first code block** will be the models of the template, i.e. the parameters of generated function.

If your template doesn't need any model input, then just leave it blank.

## Helper / Include other template

As `gorazor` compiles templates to go function, embedding another template is just calling the generated function, like any other go function.

However, if the templates are designed to be embedded, they must be under `helper` namespace, i.e. put them in `helper` sub-folder.

So, using a helper template is similar to:

```html

@if msg != "" {
	<div>@helper.ShowMsg(msg)</div>
}

```

`gorazor` won't HTML escape the output of `helper.XXX`.

Please use [example](https://github.com/sipin/gorazor/blob/master/examples/tpl/home.gohtml) for reference.

## Layout & Section

The syntax for declaring layout is a bit tricky, in the example mentioned above:

```go
@{
	import  (
		"tpl/layout"
	)

	layout := layout.Base //Use layout package's **base func** for layout
}
```

`"tpl/layout"` **is** the layout package namespace, and the `layout` variable refers to `"layout.Base"` func, which should be generated by `tpl/layout/base.gohtml`.

> Must use `layout` as the variable name

### Package / Variable convention

* The namespace/folder name for layout templates **must be** `layout`
  * `gorazor` relies on this to determine if a template is for layout
* The template `variable name` also **must be** `layout`

A layout file `tpl/layout/base.gohtml` may look like:

```html
@{
	var body string
	var sidebar string
	var footer string
	var title string
	var css string
	var js string
}

<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8" />
	<title>@title</title>
</head>
<body>
    <div class="container">@body</div>
    <div class="sidebar">@sidebar</div>
    <div class="footer">@footer</div>
	@js
  </body>
</html>
```

It's just a usual `gorazor` template, but:

* First param must be `var body string` (As it's always required, maybe we could remove it in future?)
* All params **must be** string, each param is considered as a **section**, the variable name is the **section name**.
* Under `layout` package, i.e. within "layout" folder.
  * Optionally, use `isLayout := true` to declare a template as layout

A template using such layout `tpl/index.gohtml` may look like:

```html
@{
	import (
		"tpl/layout"
	)

	layout := layout.Base
}

@section footer {
	<div>Copyright 2014</div>
}

<h2>Welcome to homepage</h2>
```

It's also possible to use import alias:

```html
@{
	import (
		share "tpl/layout"
	)

	layout := share.Base
}
```

With the page, the page content will be treated as the `body` section in the layout.

The other section content need to be wrapped with
```
@section SectionName {
	....
}
```

The template doesn't need to specify all sections defined in the layout. If a section is not specified, it will be considered as `""`.

Thus, it's possible for the layout to define default section content in such manner:

```html
@{
	var body string
	var sidebar string
}

<body>
    <div class="container">@body</div>
    @if sidebar == "" {
    <div class="sidebar">I'm the default side bar</div>
	} else {
    <div class="sidebar">@sidebar</div>
	}
</body>
```

* A layout should be able to use another layout, it's just function call.

# Conventions

* Template **folder name** will be used as **package name** in generated code
* Template file name must has the extension name `.gohtml`
* Template strip of `.gohtml` extension name will be used as the **function name** in generated code, with **first letter Capitalized**.
  * So that the function will be accessible to other modules. (I hate GO about this.)
* Helper templates **must** has the package name **helper**, i.e. in `helper` folder.
* Layout templates **must** has the package name **layout**, i.e. in `layout` folder.

# Example

Here is a simple example of [gorazor templates](https://github.com/sipin/gorazor/tree/master/examples/tpl) and the corresponding [generated codes](https://github.com/sipin/gorazor/tree/master/examples/gen).

# FAQ

## IDE / Editor support?

### Sublime Text 2/3

* **Syntax highlight** Search & install `gorazor` via Package Control.
* **Context aware auto-completion**, you may need to [manually modify](https://github.com/sipin/GoSublime/commit/fd0b979e7cc1d8f2438bb314399c2456d16f3ffb) GoSublime package, bascially replace `gscomplete.py` in with [this](https://raw.githubusercontent.com/sipin/GoSublime/gorazor/gscomplete.py) and `gslint.py` with [this](https://raw.githubusercontent.com/sipin/GoSublime/gorazor/gslint.py)

### Emacs
[web-mode](http://web-mode.org/) supports Razor template engine, so add this into your Emacs config file:

```lisp
(require 'web-mode)
(add-hook 'web-mode-hook  'my-web-mode-hook)
(add-to-list 'auto-mode-alist '("\\.html?\\'" . web-mode))
(add-to-list 'auto-mode-alist '("\\.gohtml\\'" . web-mode))
(add-to-list 'auto-mode-alist '("\\.js\\'" . web-mode))
(setq web-mode-engines-alist '(("razor" . "\\.gohtml\\'")))
```

# Credits

The very [first version](https://github.com/sipin/gorazor/releases/tag/vash) of `gorazor` is a hack of razor's port in javascript: [vash](https://github.com/kirbysayshi/vash), thus requires node's to run.

`gorazor` has been though several rounds of refactoring, and it has completely rewritten in pure Go. Nonetheless, THANK YOU [@kirbysayshi](https://github.com/kirbysayshi) for Vash! Without Vash, `gorazor` may never start.

# LICENSE

[LICENSE](LICENSE)? Well, [WTFPL](http://www.wtfpl.net/about/) and [996.icu](LICENSE).

# Todo

[Todo](https://github.com/sipin/gorazor/blob/master/todo.md)
