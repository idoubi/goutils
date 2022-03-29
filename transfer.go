package goutils

import (
	"bytes"
	"encoding/xml"
	"strconv"
	"strings"

	xml2json "github.com/basgys/goxml2json"
)

// TextEncode encode text with emoji
func TextEncode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}

	return ret
}

// TextDecode decode text with emoji
func TextDecode(raw string) {

}

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

// StrToKv transfer multi line content to kv options
func StrToKv(str string) map[string]string {
	kv := map[string]string{}

	lines := strings.Split(str, "\r\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		arr := strings.SplitN(line, ":", 2)
		if len(arr) != 2 {
			continue
		}
		kv[arr[0]] = arr[1]
	}

	return kv
}

// StrToArr transfer multi line content to array
func StrToArr(str string) []string {
	arr := []string{}

	lines := strings.Split(str, "\r\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		arr = append(arr, line)
	}

	return arr
}
