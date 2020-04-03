// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014
type DatatypeUUID struct {
	Uuid string
}

const DatatypeUUIDAvroCRC64Fingerprint = "\xfc\xa43\x98\xee\xe0p\xe2"

func NewDatatypeUUID() *DatatypeUUID {
	return &DatatypeUUID{}
}

func DeserializeDatatypeUUID(r io.Reader) (*DatatypeUUID, error) {
	t := NewDatatypeUUID()
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

func DeserializeDatatypeUUIDFromSchema(r io.Reader, schema string) (*DatatypeUUID, error) {
	t := NewDatatypeUUID()

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

func writeDatatypeUUID(r *DatatypeUUID, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Uuid, w)
	if err != nil {
		return err
	}
	return err
}

func (r *DatatypeUUID) Serialize(w io.Writer) error {
	return writeDatatypeUUID(r, w)
}

func (r *DatatypeUUID) Schema() string {
	return "{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"bodyworks.datatype.UUID\",\"type\":\"record\"}"
}

func (r *DatatypeUUID) SchemaName() string {
	return "bodyworks.datatype.UUID"
}

func (_ *DatatypeUUID) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *DatatypeUUID) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *DatatypeUUID) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *DatatypeUUID) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *DatatypeUUID) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *DatatypeUUID) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *DatatypeUUID) SetString(v string)   { panic("Unsupported operation") }
func (_ *DatatypeUUID) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatatypeUUID) Get(i int) types.Field {
	switch i {
	case 0:
		return (*types.String)(&r.Uuid)
	}
	panic("Unknown field index")
}

func (r *DatatypeUUID) SetDefault(i int) {
	switch i {
	case 0:
		r.Uuid = ""
		return
	}
	panic("Unknown field index")
}

func (_ *DatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *DatatypeUUID) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *DatatypeUUID) Finalize()                        {}

func (_ *DatatypeUUID) AvroCRC64Fingerprint() []byte {
	return []byte(DatatypeUUIDAvroCRC64Fingerprint)
}
