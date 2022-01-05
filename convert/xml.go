package convert

import (
	"bytes"

	xml2json "github.com/basgys/goxml2json"
)

// Xml2Json xml转换成json
func Xml2Json(x []byte) ([]byte, error) {
	xr := bytes.NewReader(x)
	j, err := xml2json.Convert(xr)
	if err != nil {
		return nil, err
	}

	return j.Bytes(), nil
}
