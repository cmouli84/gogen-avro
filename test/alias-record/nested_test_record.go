// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     nested.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type NestedTestRecord struct {
	OtherField *NestedRecord
}

const NestedTestRecordAvroCRC64Fingerprint = "@8B\xd3\xc9\xf5\xb5\x9b"

func NewNestedTestRecord() *NestedTestRecord {
	return &NestedTestRecord{}
}

func DeserializeNestedTestRecord(r io.Reader) (*NestedTestRecord, error) {
	t := NewNestedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeNestedTestRecordFromSchema(r io.Reader, schema string) (*NestedTestRecord, error) {
	t := NewNestedTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeNestedTestRecord(r *NestedTestRecord, w io.Writer) error {
	var err error
	err = writeNestedRecord(r.OtherField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *NestedTestRecord) Serialize(w io.Writer) error {
	return writeNestedTestRecord(r, w)
}

func (r *NestedTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"OtherField\",\"type\":{\"aliases\":[\"aliasedRecord\"],\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"NestedRecord\",\"type\":\"record\"}}],\"name\":\"NestedTestRecord\",\"type\":\"record\"}"
}

func (r *NestedTestRecord) SchemaName() string {
	return "NestedTestRecord"
}

func (_ *NestedTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *NestedTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *NestedTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *NestedTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *NestedTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *NestedTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *NestedTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *NestedTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.OtherField = NewNestedRecord()

		return r.OtherField
	}
	panic("Unknown field index")
}

func (r *NestedTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (_ *NestedTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NestedTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *NestedTestRecord) Finalize()                        {}

func (_ *NestedTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NestedTestRecordAvroCRC64Fingerprint)
}
