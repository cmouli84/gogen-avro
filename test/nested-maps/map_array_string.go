// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     nested-maps.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

func writeMapArrayString(r *MapArrayString, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeArrayString(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapArrayString struct {
	keys   []string
	values [][]string
	M      map[string][]string
}

func NewMapArrayString() *MapArrayString {
	return &MapArrayString{
		keys:   make([]string, 0),
		values: make([][]string, 0),
		M:      make(map[string][]string),
	}
}

func (_ *MapArrayString) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapArrayString) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapArrayString) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapArrayString) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapArrayString) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapArrayString) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapArrayString) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapArrayString) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapArrayString) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapArrayString) SetDefault(i int)      { panic("Unsupported operation") }
func (r *MapArrayString) Finalize() {
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapArrayString) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v []string
	v = make([]string, 0)

	r.values = append(r.values, v)
	return (*ArrayStringWrapper)(&r.values[len(r.values)-1])
}

func (_ *MapArrayString) AppendArray() types.Field { panic("Unsupported operation") }
