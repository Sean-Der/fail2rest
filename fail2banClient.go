package main

import (
	"bytes"
	"errors"
	"github.com/kisielk/og-rek"
	"net"
)

func fail2banRequest(input []string) (interface{}, error) {
	c, err := net.Dial("unix", "/var/run/fail2ban/fail2ban.sock")

	if err != nil {
		return nil, errors.New("Failed to contact fail2ban socket")
	}

	p := &bytes.Buffer{}
	ogórek.NewEncoder(p).Encode(input)
	c.Write(p.Bytes())
	c.Write([]byte("<F2B_END_COMMAND>"))

	buf := make([]byte, 0)
	tmpBuf := make([]byte, 1)
	for {
		bufRead, _ := c.Read(tmpBuf)

		if bufRead != 0 {
			buf = append(buf, tmpBuf...)
		} else {
			buf = buf[:len(buf)-17]
			break
		}

	}

	dec := ogórek.NewDecoder(bytes.NewBuffer(buf))
	v, err := dec.Decode()
	return v, err
}
