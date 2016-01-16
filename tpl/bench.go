package tpl

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	. "github.com/vycb/gorazortpl/tpl/bench"
	"strings"
)

func Bench(p *Page) string {
	var _buffer bytes.Buffer
	_buffer.WriteString(gorazor.HTMLEscape(Header(p)))
	_buffer.WriteString("\n<h1>Hello ")
	_buffer.WriteString(gorazor.HTMLEscape(strings.TrimSpace(p.Title)))
	_buffer.WriteString(" !</h1>\n\n<p>Here's a list of your favorite colors:</p>\n<ul>\n\t")
	for _, colorName := range p.FavoriteColors {

		_buffer.WriteString("<li>")
		_buffer.WriteString(gorazor.HTMLEscape(colorName))
		_buffer.WriteString("</li>")

	}
	_buffer.WriteString("\n</ul>")
	_buffer.WriteString(gorazor.HTMLEscape(Footer(p)))

	return _buffer.String()
}
