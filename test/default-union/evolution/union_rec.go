// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type UnionRec struct {
	A int32
}

const UnionRecAvroCRC64Fingerprint = "1\xf9\xae\xb7W\x80#\xf9"

func NewUnionRec() *UnionRec {
	return &UnionRec{}
}

func DeserializeUnionRec(r io.Reader) (*UnionRec, error) {
	t := NewUnionRec()
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

func DeserializeUnionRecFromSchema(r io.Reader, schema string) (*UnionRec, error) {
	t := NewUnionRec()

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

func writeUnionRec(r *UnionRec, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.A, w)
	if err != nil {
		return err
	}
	return err
}

func (r *UnionRec) Serialize(w io.Writer) error {
	return writeUnionRec(r, w)
}

func (r *UnionRec) Schema() string {
	return "{\"fields\":[{\"name\":\"a\",\"type\":\"int\"}],\"name\":\"unionRec\",\"type\":\"record\"}"
}

func (r *UnionRec) SchemaName() string {
	return "unionRec"
}

func (_ *UnionRec) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *UnionRec) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *UnionRec) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *UnionRec) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *UnionRec) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *UnionRec) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *UnionRec) SetString(v string)   { panic("Unsupported operation") }
func (_ *UnionRec) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UnionRec) Get(i int) types.Field {
	switch i {
	case 0:
		return (*types.Int)(&r.A)
	}
	panic("Unknown field index")
}

func (r *UnionRec) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (_ *UnionRec) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionRec) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionRec) Finalize()                        {}

func (_ *UnionRec) AvroCRC64Fingerprint() []byte {
	return []byte(UnionRecAvroCRC64Fingerprint)
}
