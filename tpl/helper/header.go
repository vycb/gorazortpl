package helper

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Header(u *User) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n\t<meta charset=\"UTF-8\">\n\t<title>")
	_buffer.WriteString(gorazor.HTMLEscape(u.Name))
	_buffer.WriteString(" ")
	_buffer.WriteString(gorazor.HTMLEscape(u.Email))
	_buffer.WriteString("</title>\n</head>\n<body>\n<div>Page Header</div>")

	return _buffer.String()
}
