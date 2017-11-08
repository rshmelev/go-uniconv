package serializing

import (
	"encoding/json"

	uc "github.com/rshmelev/go-uniconv"
)

//import first "github.com/rshmelev/go-uniconv/first"

// extend
type UniversalSerializingConverter interface {
	uc.UniversalConverter
	Map() map[string]interface{}
	Failed() error

	JsonStr(fallback ...string) string
	JsonIndentStr(indentation string, fallback ...string) string
	XmlStr(fallback ...string) string
	XmlIndentStr(indentation string, fallback ...string) string
	YamlStr(indentation string, fallback ...string) string

	AsJsonStruct() UniversalSerializingConverter
	AsXmlStruct() UniversalSerializingConverter
	AsYamlStruct() UniversalSerializingConverter
	As(format string) UniversalSerializingConverter

	Into(something interface{}) error
}

// extend
type SoftConverter struct {
	uc.SoftConverter
	err error

	//TreatDataAs string // can be "json"
}

// new ConverterFunc
var Cast ConverterFunc = func(a interface{}) UniversalSerializingConverter {
	x := &SoftConverter{}
	x.Value = a
	return x
}

type ConverterFunc func(interface{}) UniversalSerializingConverter

//---------------------------------------- new methods

func (c *SoftConverter) Failed() error {
	return c.err
}

//---------------------------------------- predefined serializers

func (c *SoftConverter) JsonStr(fallback ...string) string {
	return c.JsonIndentStr("  ", fallback...)
}
func (c *SoftConverter) JsonIndentStr(indentation string, fallback ...string) string {
	return SerializeToStringWithFallback(c.Value, indentation, "json", fallback...)
}
func (c *SoftConverter) XmlStr(fallback ...string) string {
	return c.XmlIndentStr("  ", fallback...)
}
func (c *SoftConverter) XmlIndentStr(indentation string, fallback ...string) string {
	return SerializeToStringWithFallback(c.Value, indentation, "xml", fallback...)
}
func (c *SoftConverter) YamlStr(indentation string, fallback ...string) string {
	return SerializeToStringWithFallback(c.Value, indentation, "yaml", fallback...)
}

func (c *SoftConverter) jsonBytes() []byte {
	var bytes []byte
	bytes, c.err = json.Marshal(c.Value)
	if bytes == nil {
		bytes, _ = json.Marshal(struct{ SerializationError string }{c.err.Error()})
	}
	return bytes
}

//----------------------------------------------------------

func (c *SoftConverter) asMap() *SoftConverter {
	if j, err := json.MarshalIndent(c.Value, "", ""); err == nil {
		c.Value = j
	} else {
		c.err = err
		c.Value = map[string]interface{}{"SerializationError": err.Error()}
	}
	return c
}

// note: requires struct, not array or single value!
func (c *SoftConverter) AsJsonStruct() UniversalSerializingConverter { return c.As("json-struct") }
func (c *SoftConverter) AsXmlStruct() UniversalSerializingConverter  { return c.As("xml-struct") }
func (c *SoftConverter) AsYamlStruct() UniversalSerializingConverter { return c.As("yaml-struct") }
func (c *SoftConverter) As(format string) UniversalSerializingConverter {
	bytes := tryToGetArrayOfBytes(c.Value)
	if bytes != nil {
		if m, err := DeserializeToStructurized(bytes, format); err == nil {
			c.Value = m
		} else {
			c.err = err
			c.Value = map[string]interface{}{"DeserializationError": err.Error()}
		}
	} else {
		c.Value = map[string]interface{}{}
	}

	return c
}

//-----------------------------------------------------

func (c *SoftConverter) Map() map[string]interface{} {
	something := c.Value
	switch n := something.(type) {
	case map[string]interface{}:
		return n
	default:
		return c.asMap().Map()
	}
}

func (c *SoftConverter) Into(something interface{}) error {
	if c.err != nil {
		return c.err
	}
	var bytes []byte = nil

	switch n := something.(type) {
	case string:
		bytes = []byte(n)
	case []byte:
		bytes = n
	default:
		bytes = c.jsonBytes()
	}

	if c.err != nil {
		return c.err
	}

	c.err = json.Unmarshal(bytes, something)

	return c.err
}
