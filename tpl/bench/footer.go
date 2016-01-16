package bench

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Footer(p *Page) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<div>")
	_buffer.WriteString(gorazor.HTMLEscape(p.Year))
	_buffer.WriteString(" Year</div>\n</body>\n</html>")

	return _buffer.String()
}
