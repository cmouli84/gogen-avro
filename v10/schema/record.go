package schema

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/actgardner/gogen-avro/v10/generator"
)

type RecordDefinition struct {
	name     QualifiedName
	aliases  []QualifiedName
	fields   []*Field
	doc      string
	metadata map[string]interface{}
}

func NewRecordDefinition(name QualifiedName, aliases []QualifiedName, fields []*Field, doc string, metadata map[string]interface{}) *RecordDefinition {
	return &RecordDefinition{
		name:     name,
		aliases:  aliases,
		fields:   fields,
		doc:      doc,
		metadata: metadata,
	}
}

func (r *RecordDefinition) AvroName() QualifiedName {
	return r.name
}

func (r *RecordDefinition) Name() string {
	return generator.ToPublicName(r.name.String())
}

func (r *RecordDefinition) Package() string {
	return getPackageName(r.name.Namespace + r.name.Name)
}

func (r *RecordDefinition) GetImportPackages() []string {
	packageMap := make(map[string]bool)
	for _, field := range r.fields {
		if field.HasDefault() {
			fieldSetters, _ := r.FullQualifiedDefaultForField(field)
			fieldSetterParts := strings.Split(fieldSetters, "\n")
			for _, setter := range fieldSetterParts {
				parts := strings.Split(setter, " = ")
				if len(parts) == 2 && strings.HasPrefix(parts[1], "com") {
					fieldType := strings.TrimSpace(parts[1])
					pkg := strings.Split(fieldType, ".")
					if len(pkg) == 2 {
						packageMap[pkg[0]] = true
					}
				}
			}
		}
		packageMap[field.Package()] = true
	}
	packages := make([]string, 0, len(packageMap))
	for pkg := range packageMap {
		if pkg != "" {
			if pkg == r.Package() {
				continue
			}

			if pkg != generator.PackageName {
				pkg = fmt.Sprintf("%s/%s", generator.PackageName, pkg)
			}
			if generator.BasePackageName == "" {
				packages = append(packages, pkg)
			} else {
				packages = append(packages, fmt.Sprintf("%s/%s", generator.BasePackageName, pkg))
			}
		}
	}
	sort.Strings(packages)
	return packages
}

func (r *RecordDefinition) GoType() string {
	return r.Name()
}

func (r *RecordDefinition) FullQualifiedGoType() string {
	return fmt.Sprintf("%s.%s", r.Package(), r.GoType())
}

func (r *RecordDefinition) Aliases() []QualifiedName {
	return r.aliases
}

func (r *RecordDefinition) SerializerMethod() string {
	return fmt.Sprintf("Write%v", r.Name())
}

func (r *RecordDefinition) FullQualifiedSerializerMethod() string {
	return fmt.Sprintf("%s.Write%v", r.Package(), r.Name())
}

func (r *RecordDefinition) NewWriterMethod() string {
	return fmt.Sprintf("New%vWriter", r.Name())
}

func (s *RecordDefinition) Attribute(name string) interface{} {
	return s.metadata[name]
}

func (r *RecordDefinition) Definition(scope map[QualifiedName]interface{}) (interface{}, error) {
	if _, ok := scope[r.name]; ok {
		return r.name.String(), nil
	}
	metadata := copyDefinition(r.metadata)
	scope[r.name] = 1
	fields := make([]map[string]interface{}, 0)
	for _, f := range r.fields {
		def, err := f.Definition(scope)
		if err != nil {
			return nil, err
		}
		fields = append(fields, def)
	}

	metadata["fields"] = fields
	return metadata, nil
}

func (r *RecordDefinition) ConstructorMethod() string {
	return fmt.Sprintf("New%v()", r.Name())
}

func (r *RecordDefinition) FullQualifiedConstructorMethod() string {
	return fmt.Sprintf("%s.New%v()", r.Package(), r.Name())
}

func (r *RecordDefinition) DefaultForField(f *Field) (string, error) {
	result, err := f.Type().DefaultValue(fmt.Sprintf("r.%v", f.GoName()), f.Default())
	return result, err
}

func (r *RecordDefinition) FullQualifiedDefaultForField(f *Field) (string, error) {
	result, err := f.Type().FullQualifiedDefaultValue(fmt.Sprintf("r.%v", f.GoName()), f.Default())
	return result, err
}

func (r *RecordDefinition) ConstructableForField(f *Field) string {
	if constructor, ok := getConstructableForType(f.Type()); ok {
		return fmt.Sprintf("r.%v = %v\n", f.GoName(), constructor.ConstructorMethod())
	}
	return ""
}

func (r *RecordDefinition) FullQualifiedConstructableForField(f *Field) string {
	if constructor, ok := getConstructableForType(f.Type()); ok {
		return fmt.Sprintf("r.%v = %v\n", f.GoName(), constructor.FullQualifiedConstructorMethod())
	}
	return ""
}

func (r *RecordDefinition) RecordReaderTypeName() string {
	return r.Name() + "Reader"
}

// FieldByName finds a field in the reader schema whose name or aliases match a name in the writer schema.
func (r *RecordDefinition) FieldByName(field string) *Field {
	for _, f := range r.fields {
		if f.NameMatchesAliases(field) {
			return f
		}
	}
	return nil
}

func (r *RecordDefinition) DefaultValue(lvalue string, rvalue interface{}) (string, error) {
	items := rvalue.(map[string]interface{})
	fieldSetters := ""
	for k, v := range items {
		field := r.FieldByName(k)
		fieldSetter, err := field.Type().DefaultValue(fmt.Sprintf("%v.%v", lvalue, field.GoName()), v)
		if err != nil {
			return "", err
		}

		fieldSetters += fieldSetter + "\n"
	}
	return fieldSetters, nil
}

func (r *RecordDefinition) FullQualifiedDefaultValue(lvalue string, rvalue interface{}) (string, error) {
	items := rvalue.(map[string]interface{})
	fieldSetters := ""
	for k, v := range items {
		field := r.FieldByName(k)
		fieldSetter, err := field.Type().FullQualifiedDefaultValue(fmt.Sprintf("%v.%v", lvalue, field.GoName()), v)
		if err != nil {
			return "", err
		}

		fieldSetters += fieldSetter + "\n"
	}
	return fieldSetters, nil
}

func (r *RecordDefinition) Fields() []*Field {
	return r.fields
}

func (s *RecordDefinition) IsReadableBy(d Definition) bool {
	_, ok := d.(*RecordDefinition)
	return ok && hasMatchingName(s.AvroName(), d)
}

func (s *RecordDefinition) WrapperType() string {
	return "types.Record"
}

func (s *RecordDefinition) FullQualifiedWrapperType() string {
	return s.WrapperType()
}

func (s *RecordDefinition) WrapperPointer() bool {
	return false
}

func (s *RecordDefinition) Doc() string {
	return strings.ReplaceAll(s.doc, "\n", " ")
}

func (s *RecordDefinition) Schema() (string, error) {
	def0, err := s.Definition(make(map[QualifiedName]interface{}))
	if err != nil {
		return "", err
	}
	def := def0.(map[string]interface{})
	delete(def, "namespace")
	def["name"] = s.name.String()
	jsonBytes, err := json.Marshal(def)
	return string(jsonBytes), err
}

func (s *RecordDefinition) Children() []AvroType {
	children := make([]AvroType, len(s.fields))
	for i, field := range s.fields {
		children[i] = field.Type()
	}
	return children
}

func (s *RecordDefinition) GetReference() bool {
	return true
}
