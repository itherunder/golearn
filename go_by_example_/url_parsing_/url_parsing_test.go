package url_parsing_

import (
	"fmt"
	"net"
	"net/url"
	"testing"
)

var pln = fmt.Println

func TestUrlParsing(t *testing.T) {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	pln(u.User)
	pln(u.User.Username())
	p, _ := u.User.Password()
	pln(p)

	pln(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	pln(host)
	pln(port)

	pln(u.Path)
	pln(u.Fragment)

	pln(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	pln(m)
	pln(m["k"][0])
}
