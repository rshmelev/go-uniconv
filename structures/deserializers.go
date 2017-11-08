package serializing

import (
	"encoding/json"
	"encoding/xml"
	"github.com/ghodss/yaml"
)

type Deserializer func(data []byte, dest ...map[string]interface{}) (map[string]interface{}, error)

var Deserializers = map[string]Deserializer{
	"json-struct": func(bytes []byte, dest ...map[string]interface{}) (map[string]interface{}, error) {
		destmap := mapPlease(dest...)
		e := json.Unmarshal(bytes, &destmap)
		return destmap, e
	},
	"xml-struct": func(bytes []byte, dest ...map[string]interface{}) (map[string]interface{}, error) {
		destmap := mapPlease(dest...)
		e := xml.Unmarshal(bytes, &destmap)
		return destmap, e
	},
	"yaml-struct": func(bytes []byte, dest ...map[string]interface{}) (map[string]interface{}, error) {
		destmap := mapPlease(dest...)
		e := yaml.Unmarshal(bytes, &destmap)
		return destmap, e
	},
}

func mapPlease(m ...map[string]interface{}) map[string]interface{} {
	var w map[string]interface{}
	if len(m) > 0 && m[0] != nil {
		w = m[0]
	} else {
		w = map[string]interface{}{}
	}
	return w
}

func DeserializeToStructurized(data []byte, format string, dest ...map[string]interface{}) (map[string]interface{}, error) {

	var w = mapPlease(dest...)

	deserializer, ok := Deserializers[format]
	if !ok {
		deserializer = Deserializers["json-struct"]
	}

	_, e := deserializer(data, w)
	return w, e
}

func tryToGetArrayOfBytes(something interface{}) []byte {
	var bytes []byte
	if something != nil {
		switch n := something.(type) {
		case string:
			bytes = []byte(n)
		case []byte:
			bytes = n
		default:
			return nil
		}
	}
	return bytes
}
