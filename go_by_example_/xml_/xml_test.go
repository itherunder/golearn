package xml_

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/yezihack/colorlog"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func TestXml(t *testing.T) {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	colorlog.Info(string(out))

	colorlog.Info(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		colorlog.Error(err.Error())
		panic(err)
	}
	colorlog.Info(p.String())

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plant   []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plant = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	colorlog.Info(string(out))
}
