package serializing

import (
	"encoding/json"
	"encoding/xml"
	"github.com/ghodss/yaml"
)

type Serializer func(something interface{}, param string) ([]byte, error)

var Serializers = map[string]Serializer{
	"json": func(something interface{}, param string) ([]byte, error) {
		return json.MarshalIndent(something, "", param)
	},
	"xml": func(something interface{}, param string) ([]byte, error) {
		return xml.MarshalIndent(something, "", param)
	},
	"yaml": func(something interface{}, param string) ([]byte, error) {
		return yaml.Marshal(something)
	},
}

func SerializeToStringWithFallback(something interface{}, param string, way string, fallback ...string) string {
	b, e := Serialize(something, param, way)
	if e != nil || b == nil {
		if len(fallback) == 0 {
			return Cast(struct{ SerializationError string }{e.Error()}).Str()
		}
		return fallback[0]
	}
	return string(b)
}
func Serialize(something interface{}, param string, way string) ([]byte, error) {
	js, ok := Serializers[way]
	if !ok {
		js, _ = Serializers["json"] // hm... should i probably panic here?
	}
	return js(something, param)
}
