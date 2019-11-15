package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/magiconair/properties"
)

func PropertiesToJSON(sources map[string]string) (ret string, err error) {
	pro := properties.NewProperties()
	for k, v := range sources {
		_, ok, _ := pro.Set(k, v)
		if !ok {
			return "", fmt.Errorf("[PropertiesToJSON] convert sources ro properties error. k:%s; v:%s", k, v)
		}
	}
	buf := new(bytes.Buffer)
	if err = pro.Decode(&buf); err != nil {
		return "", fmt.Errorf("[PropertiesToJSON] decode properties error: %s", err)
	}

	out, err := jsonMarshal(buf)
	if err != nil {
		return "", fmt.Errorf("[PropertiesToJSON] convert properties to json error: %s", err)
	}

	return string(out), nil
}

func jsonMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
