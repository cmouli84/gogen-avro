package schema

import (
	"fmt"
)

// Reference to a user-defined type (fixed, enum, record). Just a wrapper with a name around a Definition.
type Reference struct {
	TypeName QualifiedName
	Def      Definition
}

func NewReference(typeName QualifiedName) *Reference {
	return &Reference{
		TypeName: typeName,
	}
}

func (s *Reference) Name() string {
	return s.Def.Name()
}

func (s *Reference) Package() string {
	return s.Def.Package()
}

func (s *Reference) GoType() string {
	return s.Def.GoType()
}

func (s *Reference) FullQualifiedGoType() string {
	if _, ok := s.Def.(PrimitiveType); ok {
		return s.GoType()
	}

	return fmt.Sprintf("%s.%s", s.Def.Package(), s.GoType())
}

func (s *Reference) SerializerMethod() string {
	return s.Def.SerializerMethod()
}

func (s *Reference) FullQualifiedSerializerMethod() string {
	if _, ok := s.Def.(PrimitiveType); ok {
		return s.SerializerMethod()
	}
	return fmt.Sprintf("%s.%s", s.Def.Package(), s.SerializerMethod())
}

func (s *Reference) Attribute(name string) interface{} {
	return s.Def.Attribute(name)
}

func (s *Reference) Definition(scope map[QualifiedName]interface{}) (interface{}, error) {
	return s.Def.Definition(scope)
}

func (s *Reference) DefaultValue(lvalue string, rvalue interface{}) (string, error) {
	return s.Def.DefaultValue(lvalue, rvalue)
}

func (s *Reference) FullQualifiedDefaultValue(lvalue string, rvalue interface{}) (string, error) {
	return s.Def.FullQualifiedDefaultValue(lvalue, rvalue)
}

func (s *Reference) WrapperType() string {
	return s.Def.WrapperType()
}

func (s *Reference) FullQualifiedWrapperType() string {
	if _, ok := s.Def.(PrimitiveType); ok {
		return s.WrapperType()
	}
	if _, ok := s.Def.(*RecordDefinition); ok {
		return s.WrapperType()
	}
	return fmt.Sprintf("%s.%s", s.Package(), s.WrapperType())
}

func (s *Reference) WrapperPointer() bool {
	return false
}

func (s *Reference) IsReadableBy(f AvroType) bool {
	if union, ok := f.(*UnionField); ok {
		for _, t := range union.AvroTypes() {
			if s.IsReadableBy(t) {
				return true
			}
		}
	}
	if reader, ok := f.(*Reference); ok {
		return s.Def.IsReadableBy(reader.Def)
	}
	return false
}

func (s *Reference) Children() []AvroType {
	// References can only point to Definitions and thus have no children
	return []AvroType{}
}

func (s *Reference) UnionKey() string {
	return s.Def.AvroName().String()
}
