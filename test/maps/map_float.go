// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
package avro

import (
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

func writeMapFloat(r *MapFloat, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteFloat(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapFloat struct {
	keys   []string
	values []float32
	M      map[string]float32
}

func NewMapFloat() *MapFloat {
	return &MapFloat{
		keys:   make([]string, 0),
		values: make([]float32, 0),
		M:      make(map[string]float32),
	}
}

func (_ *MapFloat) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapFloat) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapFloat) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapFloat) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapFloat) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapFloat) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapFloat) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapFloat) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapFloat) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapFloat) SetDefault(i int)      { panic("Unsupported operation") }
func (r *MapFloat) Finalize() {
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapFloat) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v float32
	r.values = append(r.values, v)
	return (*types.Float)(&r.values[len(r.values)-1])
}

func (_ *MapFloat) AppendArray() types.Field { panic("Unsupported operation") }
