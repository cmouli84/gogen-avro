package schema

import (
	"strings"

	"github.com/actgardner/gogen-avro/v10/generator"
)

type AvroType interface {
	Name() string
	Package() string
	GoType() string
	FullQualifiedGoType() string
	// The key to use in JSON-encoding a union with this value
	UnionKey() string

	// The name of the method which writes this field onto the wire
	SerializerMethod() string
	FullQualifiedSerializerMethod() string

	Children() []AvroType

	Attribute(name string) interface{}
	Definition(scope map[QualifiedName]interface{}) (interface{}, error)
	DefaultValue(lvalue string, rvalue interface{}) (string, error)
	FullQualifiedDefaultValue(lvalue string, rvalue interface{}) (string, error)

	// WrapperType is the VM type to wrap this value in, if applicable
	WrapperType() string
	FullQualifiedWrapperType() string
	// WrapperPointer is whether the VM wrapper type needs to be a pointer
	WrapperPointer() bool
	IsReadableBy(f AvroType) bool
}

func getPackageName(namespace string) string {
	if namespace == "" {
		return generator.PackageName
	}
	return strings.ToLower(strings.ReplaceAll(namespace, ".", ""))
}
