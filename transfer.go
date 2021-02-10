package goutils

import (
	"bytes"
	"encoding/xml"
	"strings"

	xml2json "github.com/basgys/goxml2json"
)

// XML2JSON transfer xml struct to json object
func XML2JSON(x []byte) ([]byte, error) {
	xr := bytes.NewReader(x)
	j, err := xml2json.Convert(xr)
	if err != nil {
		return nil, err
	}

	return j.Bytes(), nil
}

// Map2XML transfer map struct to xml struct
func Map2XML(m map[string]string, opts ...interface{}) ([]byte, error) {
	rootTag := "xml"
	if len(opts) > 0 {
		// the first opts is the root tag of xml struct
		if v, ok := opts[0].(string); ok {
			rootTag = v
		}
	}

	d := mdata{xml.Name{Local: rootTag}, m}

	bt, err := xml.Marshal(d)
	if err != nil {
		return nil, err
	}

	return bt, nil
}

type mdata struct {
	XMLName xml.Name
	data    map[string]string
}

// MarshalXML xml encode
func (m mdata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m.data) == 0 {
		return nil
	}

	start.Name.Local = m.XMLName.Local

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m.data {
		if strings.HasPrefix(v, "cdata:") {
			v = strings.Replace(v, "cdata:", "", 1)
			xs := struct {
				XMLName xml.Name
				Value   interface{} `xml:",cdata"`
			}{xml.Name{Local: k}, v}
			e.Encode(xs)
		} else {
			xs := struct {
				XMLName xml.Name
				Value   interface{} `xml:",chardata"`
			}{xml.Name{Local: k}, v}
			e.Encode(xs)
		}
	}

	return e.EncodeToken(start.End())
}
