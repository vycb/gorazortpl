package bench

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Header(p *Page) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n\t<meta charset=\"UTF-8\">\n\t<title>")
	_buffer.WriteString(gorazor.HTMLEscape(p.Title))
	_buffer.WriteString("</title>\n</head>\n<body>")

	return _buffer.String()
}
