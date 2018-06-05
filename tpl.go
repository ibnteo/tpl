// Template engine similar to sprintf
package tpl

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"net/url"
	"os"
	"strings"
)

// %e = html.EscapeString(s)
// %q = url.QueryEscape(s)
// %s = string
// %n = text<br>
// %% = %
func Write(w io.Writer, format string, a ...interface{}) (n int64, err error) {
	b := new(bytes.Buffer)
	doTemplate(b, format, a)
	n, err = b.WriteTo(w)
	return
}

func Print(format string, a ...interface{}) (n int64, err error) {
	return Write(os.Stdout, format, a...)
}

func Format(format string, a ...interface{}) string {
	b := new(bytes.Buffer)
	doTemplate(b, format, a)
	return b.String()
}

func doTemplate(b *bytes.Buffer, format string, a []interface{}) {
	end := len(format)
	argNum := 0
	for i := 0; i < end; i++ {
		if format[i] != '%' {
			b.WriteByte(format[i])
			continue
		}
		i++
		if i >= end {
			b.WriteByte(format[i-1])
			continue
		}
		if format[i] == '%' {
			b.WriteByte(format[i])
		} else if format[i] == 'e' || format[i] == 's' || format[i] == 'q' || format[i] == 'n' {
			if argNum >= len(a) {
				b.WriteString(format[i-1:i+1] + "(MISSING)")
			} else {
				argString := fmt.Sprintf("%v", a[argNum])
				switch format[i] {
				case 'e':
					b.WriteString(html.EscapeString(argString))
				case 's':
					b.WriteString(argString)
				case 'q':
					b.WriteString(url.QueryEscape(argString))
				case 'n':
					b.WriteString(strings.Replace(strings.Replace(html.EscapeString(argString), "\n", "<br>", -1), "\r", "", -1))
				}
			}
			argNum++
		} else {
			b.WriteString(format[i-1 : i+1])
		}
	}
	if argNum < len(a) {
		b.WriteString("(BADINDEX)")
	}
}
