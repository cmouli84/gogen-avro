// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     primitives.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

// primitive record
type PrimitiveTestRecord struct {
	// int field
	IntField int32 `json:"IntField"`

	LongField int64 `json:"LongField"`

	FloatField float32 `json:"FloatField"`

	DoubleField float64 `json:"DoubleField"`

	StringField string `json:"StringField"`

	BoolField bool `json:"BoolField"`

	BytesField []byte `json:"BytesField"`
}

const PrimitiveTestRecordAvroCRC64Fingerprint = "۲\x16\xe9\xfet@\x13"

func NewPrimitiveTestRecord() *PrimitiveTestRecord {
	return &PrimitiveTestRecord{}
}

func DeserializePrimitiveTestRecord(r io.Reader) (*PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()
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

func DeserializePrimitiveTestRecordFromSchema(r io.Reader, schema string) (*PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()

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

func writePrimitiveTestRecord(r *PrimitiveTestRecord, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.LongField, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.FloatField, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.DoubleField, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *PrimitiveTestRecord) Serialize(w io.Writer) error {
	return writePrimitiveTestRecord(r, w)
}

func (r *PrimitiveTestRecord) Schema() string {
	return "{\"doc\":\"primitive\\nrecord\",\"fields\":[{\"doc\":\"int\\nfield\",\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"DoubleField\",\"type\":\"double\"},{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"PrimitiveTestRecord\",\"type\":\"record\"}"
}

func (r *PrimitiveTestRecord) SchemaName() string {
	return "PrimitiveTestRecord"
}

func (_ *PrimitiveTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PrimitiveTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.IntField}
	case 1:
		return &types.Long{Target: &r.LongField}
	case 2:
		return &types.Float{Target: &r.FloatField}
	case 3:
		return &types.Double{Target: &r.DoubleField}
	case 4:
		return &types.String{Target: &r.StringField}
	case 5:
		return &types.Boolean{Target: &r.BoolField}
	case 6:
		return &types.Bytes{Target: &r.BytesField}
	}
	panic("Unknown field index")
}

func (r *PrimitiveTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PrimitiveTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *PrimitiveTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) Finalize()                        {}

func (_ *PrimitiveTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(PrimitiveTestRecordAvroCRC64Fingerprint)
}
