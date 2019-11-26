package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/magiconair/properties"
)

func PropertiesToJSON(sources map[string]string) (ret string, err error) {
	pro := properties.LoadMap(sources)
	buf := new(bytes.Buffer)
	if _, err := pro.Write(buf, properties.UTF8); err != nil {
		return "", fmt.Errorf("[PropertiesToJSON] decode properties error: %s", err)
	}

	out, err := jsonMarshal(buf.Bytes())
	if err != nil {
		return "", fmt.Errorf("[PropertiesToJSON] convert properties to json error: %s", err)
	}
	var data interface{}
	json.Unmarshal(out, &data)
	fmt.Println(data)
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

func DotStringToJSON(key, value string, obj interface{}) string {
	parts := strings.Split(key, ".")
	if len(parts) == 1 {
		return fmt.Sprintf("{\"%s\":\"%s\"}", key, value)
	}

	last := parts[len(parts)-1]
	parts = parts[:len(parts)-2]
	for part := parts[0]; len(parts) > 0; {
		if _, ok := obj.(map[string]interface{})[part]; !ok {
			obj.(map[string]interface{})[part] = map[string]interface{}{}
		}
		obj = obj.(map[string]interface{})[part]
		if len(parts) > 1 {
			parts = parts[1:]
			part = parts[0]
		} else {
			parts = nil
		}
	}

	obj.(map[string]interface{})[last] = value

	bs, _ := json.Marshal(obj)
	return string(bs)
}

func MergeJSONs(inputs []string) string {
	_ = len(inputs)
	return ""
}
