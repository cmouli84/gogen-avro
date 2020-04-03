// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     arrays.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)

func writeArrayLong(r []int64, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteLong(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayLongWrapper []int64

func (_ *ArrayLongWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayLongWrapper) Finalize()                        {}
func (_ *ArrayLongWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r *ArrayLongWrapper) AppendArray() types.Field {
	var v int64

	*r = append(*r, v)
	return (*types.Long)(&(*r)[len(*r)-1])
}
