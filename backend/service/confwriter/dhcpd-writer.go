package confwriter

import (
	"bytes"
	"os"
	"text/template"
)

type DhcpdWriter struct {
	template *template.Template
}

func New() DhcpdWriter {
	t := template.Must(template.ParseFiles("dhcpd.tmpl"))

	return DhcpdWriter{
		template: t,
	}
}

const (
	charsPerMac = 19
	staticChars = 12
)

func (dw DhcpdWriter) WhitelistMacs(macs []string) error {
	bufferSize := staticChars + charsPerMac*len(macs)
	b := bytes.NewBuffer(make([]byte, 0, bufferSize))
	err := dw.template.Execute(b, map[any]any{"macs": macs})
	if err != nil {
		return err
	}

	f, err := os.OpenFile("dhcpd.conf", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	// FIXME risky
	_, err = f.Write(b.Bytes())
	if err != nil {
		return err
	}

	return f.Close()
}
