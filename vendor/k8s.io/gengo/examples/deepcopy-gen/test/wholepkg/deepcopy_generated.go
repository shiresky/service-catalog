// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package wholepkg

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_wholepkg_ManualStruct, InType: reflect.TypeOf(&ManualStruct{})},
		{Fn: DeepCopy_wholepkg_ManualStruct_Alias, InType: reflect.TypeOf(&ManualStruct_Alias{})},
		{Fn: DeepCopy_wholepkg_Struct_B, InType: reflect.TypeOf(&Struct_B{})},
		{Fn: DeepCopy_wholepkg_Struct_Embed_Int, InType: reflect.TypeOf(&Struct_Embed_Int{})},
		{Fn: DeepCopy_wholepkg_Struct_Embed_ManualStruct, InType: reflect.TypeOf(&Struct_Embed_ManualStruct{})},
		{Fn: DeepCopy_wholepkg_Struct_Embed_Pointer, InType: reflect.TypeOf(&Struct_Embed_Pointer{})},
		{Fn: DeepCopy_wholepkg_Struct_Embed_Struct_PrimitivePointers, InType: reflect.TypeOf(&Struct_Embed_Struct_PrimitivePointers{})},
		{Fn: DeepCopy_wholepkg_Struct_Embed_Struct_Primitives, InType: reflect.TypeOf(&Struct_Embed_Struct_Primitives{})},
		{Fn: DeepCopy_wholepkg_Struct_Embed_Struct_Slices, InType: reflect.TypeOf(&Struct_Embed_Struct_Slices{})},
		{Fn: DeepCopy_wholepkg_Struct_Empty, InType: reflect.TypeOf(&Struct_Empty{})},
		{Fn: DeepCopy_wholepkg_Struct_Everything, InType: reflect.TypeOf(&Struct_Everything{})},
		{Fn: DeepCopy_wholepkg_Struct_PrimitivePointers, InType: reflect.TypeOf(&Struct_PrimitivePointers{})},
		{Fn: DeepCopy_wholepkg_Struct_PrimitivePointers_Alias, InType: reflect.TypeOf(&Struct_PrimitivePointers_Alias{})},
		{Fn: DeepCopy_wholepkg_Struct_Primitives, InType: reflect.TypeOf(&Struct_Primitives{})},
		{Fn: DeepCopy_wholepkg_Struct_Primitives_Alias, InType: reflect.TypeOf(&Struct_Primitives_Alias{})},
		{Fn: DeepCopy_wholepkg_Struct_Slices, InType: reflect.TypeOf(&Struct_Slices{})},
		{Fn: DeepCopy_wholepkg_Struct_Slices_Alias, InType: reflect.TypeOf(&Struct_Slices_Alias{})},
		{Fn: DeepCopy_wholepkg_Struct_Struct_PrimitivePointers, InType: reflect.TypeOf(&Struct_Struct_PrimitivePointers{})},
		{Fn: DeepCopy_wholepkg_Struct_Struct_Primitives, InType: reflect.TypeOf(&Struct_Struct_Primitives{})},
		{Fn: DeepCopy_wholepkg_Struct_Struct_Slices, InType: reflect.TypeOf(&Struct_Struct_Slices{})},
	}
}

func DeepCopy_wholepkg_ManualStruct(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ManualStruct)
		out := out.(*ManualStruct)
		*out = in.DeepCopy()
		return nil
	}
}

func DeepCopy_wholepkg_ManualStruct_Alias(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ManualStruct_Alias)
		out := out.(*ManualStruct_Alias)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_B(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_B)
		out := out.(*Struct_B)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Embed_Int(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Embed_Int)
		out := out.(*Struct_Embed_Int)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Embed_ManualStruct(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Embed_ManualStruct)
		out := out.(*Struct_Embed_ManualStruct)
		*out = *in
		out.ManualStruct = in.ManualStruct.DeepCopy()
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Embed_Pointer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Embed_Pointer)
		out := out.(*Struct_Embed_Pointer)
		*out = *in
		if in.int != nil {
			in, out := &in.int, &out.int
			*out = new(int)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Embed_Struct_PrimitivePointers(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Embed_Struct_PrimitivePointers)
		out := out.(*Struct_Embed_Struct_PrimitivePointers)
		*out = *in
		if err := DeepCopy_wholepkg_Struct_PrimitivePointers(&in.Struct_PrimitivePointers, &out.Struct_PrimitivePointers, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Embed_Struct_Primitives(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Embed_Struct_Primitives)
		out := out.(*Struct_Embed_Struct_Primitives)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Embed_Struct_Slices(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Embed_Struct_Slices)
		out := out.(*Struct_Embed_Struct_Slices)
		*out = *in
		if err := DeepCopy_wholepkg_Struct_Slices(&in.Struct_Slices, &out.Struct_Slices, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Empty(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Empty)
		out := out.(*Struct_Empty)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Everything(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Everything)
		out := out.(*Struct_Everything)
		*out = *in
		out.ManualStructField = in.ManualStructField.DeepCopy()
		if in.BoolPtrField != nil {
			in, out := &in.BoolPtrField, &out.BoolPtrField
			*out = new(bool)
			**out = **in
		}
		if in.IntPtrField != nil {
			in, out := &in.IntPtrField, &out.IntPtrField
			*out = new(int)
			**out = **in
		}
		if in.StringPtrField != nil {
			in, out := &in.StringPtrField, &out.StringPtrField
			*out = new(string)
			**out = **in
		}
		if in.FloatPtrField != nil {
			in, out := &in.FloatPtrField, &out.FloatPtrField
			*out = new(float64)
			**out = **in
		}
		if err := DeepCopy_wholepkg_Struct_PrimitivePointers(&in.PrimitivePointersField, &out.PrimitivePointersField, c); err != nil {
			return err
		}
		if in.ManualStructPtrField != nil {
			in, out := &in.ManualStructPtrField, &out.ManualStructPtrField
			*out = new(ManualStruct)
			**out = (*in).DeepCopy()
		}
		if in.ManualStructAliasPtrField != nil {
			in, out := &in.ManualStructAliasPtrField, &out.ManualStructAliasPtrField
			*out = new(ManualStruct_Alias)
			**out = **in
		}
		if in.SliceBoolField != nil {
			in, out := &in.SliceBoolField, &out.SliceBoolField
			*out = make([]bool, len(*in))
			copy(*out, *in)
		}
		if in.SliceByteField != nil {
			in, out := &in.SliceByteField, &out.SliceByteField
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		if in.SliceIntField != nil {
			in, out := &in.SliceIntField, &out.SliceIntField
			*out = make([]int, len(*in))
			copy(*out, *in)
		}
		if in.SliceStringField != nil {
			in, out := &in.SliceStringField, &out.SliceStringField
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.SliceFloatField != nil {
			in, out := &in.SliceFloatField, &out.SliceFloatField
			*out = make([]float64, len(*in))
			copy(*out, *in)
		}
		if err := DeepCopy_wholepkg_Struct_Slices(&in.SlicesField, &out.SlicesField, c); err != nil {
			return err
		}
		if in.SliceManualStructField != nil {
			in, out := &in.SliceManualStructField, &out.SliceManualStructField
			*out = make([]ManualStruct, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i].DeepCopy()
			}
		}
		if in.ManualSliceField != nil {
			out.ManualSliceField = in.ManualSliceField.DeepCopy()
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_PrimitivePointers(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_PrimitivePointers)
		out := out.(*Struct_PrimitivePointers)
		*out = *in
		if in.BoolPtrField != nil {
			in, out := &in.BoolPtrField, &out.BoolPtrField
			*out = new(bool)
			**out = **in
		}
		if in.IntPtrField != nil {
			in, out := &in.IntPtrField, &out.IntPtrField
			*out = new(int)
			**out = **in
		}
		if in.StringPtrField != nil {
			in, out := &in.StringPtrField, &out.StringPtrField
			*out = new(string)
			**out = **in
		}
		if in.FloatPtrField != nil {
			in, out := &in.FloatPtrField, &out.FloatPtrField
			*out = new(float64)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_PrimitivePointers_Alias(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_PrimitivePointers_Alias)
		out := out.(*Struct_PrimitivePointers_Alias)
		*out = *in
		if in.BoolPtrField != nil {
			in, out := &in.BoolPtrField, &out.BoolPtrField
			*out = new(bool)
			**out = **in
		}
		if in.IntPtrField != nil {
			in, out := &in.IntPtrField, &out.IntPtrField
			*out = new(int)
			**out = **in
		}
		if in.StringPtrField != nil {
			in, out := &in.StringPtrField, &out.StringPtrField
			*out = new(string)
			**out = **in
		}
		if in.FloatPtrField != nil {
			in, out := &in.FloatPtrField, &out.FloatPtrField
			*out = new(float64)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Primitives(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Primitives)
		out := out.(*Struct_Primitives)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Primitives_Alias(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Primitives_Alias)
		out := out.(*Struct_Primitives_Alias)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Slices(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Slices)
		out := out.(*Struct_Slices)
		*out = *in
		if in.SliceBoolField != nil {
			in, out := &in.SliceBoolField, &out.SliceBoolField
			*out = make([]bool, len(*in))
			copy(*out, *in)
		}
		if in.SliceByteField != nil {
			in, out := &in.SliceByteField, &out.SliceByteField
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		if in.SliceIntField != nil {
			in, out := &in.SliceIntField, &out.SliceIntField
			*out = make([]int, len(*in))
			copy(*out, *in)
		}
		if in.SliceStringField != nil {
			in, out := &in.SliceStringField, &out.SliceStringField
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.SliceFloatField != nil {
			in, out := &in.SliceFloatField, &out.SliceFloatField
			*out = make([]float64, len(*in))
			copy(*out, *in)
		}
		if in.SliceStructPrimitivesField != nil {
			in, out := &in.SliceStructPrimitivesField, &out.SliceStructPrimitivesField
			*out = make([]Struct_Primitives, len(*in))
			copy(*out, *in)
		}
		if in.SliceStructPrimitivesAliasField != nil {
			in, out := &in.SliceStructPrimitivesAliasField, &out.SliceStructPrimitivesAliasField
			*out = make([]Struct_Primitives_Alias, len(*in))
			copy(*out, *in)
		}
		if in.SliceStructPrimitivePointersField != nil {
			in, out := &in.SliceStructPrimitivePointersField, &out.SliceStructPrimitivePointersField
			*out = make([]Struct_PrimitivePointers, len(*in))
			for i := range *in {
				if err := DeepCopy_wholepkg_Struct_PrimitivePointers(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.SliceStructPrimitivePointersAliasField != nil {
			in, out := &in.SliceStructPrimitivePointersAliasField, &out.SliceStructPrimitivePointersAliasField
			*out = make([]Struct_PrimitivePointers_Alias, len(*in))
			for i := range *in {
				if err := DeepCopy_wholepkg_Struct_PrimitivePointers_Alias(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.SliceSliceIntField != nil {
			in, out := &in.SliceSliceIntField, &out.SliceSliceIntField
			*out = make([][]int, len(*in))
			for i := range *in {
				if (*in)[i] != nil {
					in, out := &(*in)[i], &(*out)[i]
					*out = make([]int, len(*in))
					copy(*out, *in)
				}
			}
		}
		if in.SliceManualStructField != nil {
			in, out := &in.SliceManualStructField, &out.SliceManualStructField
			*out = make([]ManualStruct, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i].DeepCopy()
			}
		}
		if in.ManualSliceField != nil {
			out.ManualSliceField = in.ManualSliceField.DeepCopy()
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Slices_Alias(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Slices_Alias)
		out := out.(*Struct_Slices_Alias)
		*out = *in
		if in.SliceBoolField != nil {
			in, out := &in.SliceBoolField, &out.SliceBoolField
			*out = make([]bool, len(*in))
			copy(*out, *in)
		}
		if in.SliceByteField != nil {
			in, out := &in.SliceByteField, &out.SliceByteField
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		if in.SliceIntField != nil {
			in, out := &in.SliceIntField, &out.SliceIntField
			*out = make([]int, len(*in))
			copy(*out, *in)
		}
		if in.SliceStringField != nil {
			in, out := &in.SliceStringField, &out.SliceStringField
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.SliceFloatField != nil {
			in, out := &in.SliceFloatField, &out.SliceFloatField
			*out = make([]float64, len(*in))
			copy(*out, *in)
		}
		if in.SliceStructPrimitivesField != nil {
			in, out := &in.SliceStructPrimitivesField, &out.SliceStructPrimitivesField
			*out = make([]Struct_Primitives, len(*in))
			copy(*out, *in)
		}
		if in.SliceStructPrimitivesAliasField != nil {
			in, out := &in.SliceStructPrimitivesAliasField, &out.SliceStructPrimitivesAliasField
			*out = make([]Struct_Primitives_Alias, len(*in))
			copy(*out, *in)
		}
		if in.SliceStructPrimitivePointersField != nil {
			in, out := &in.SliceStructPrimitivePointersField, &out.SliceStructPrimitivePointersField
			*out = make([]Struct_PrimitivePointers, len(*in))
			for i := range *in {
				if err := DeepCopy_wholepkg_Struct_PrimitivePointers(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.SliceStructPrimitivePointersAliasField != nil {
			in, out := &in.SliceStructPrimitivePointersAliasField, &out.SliceStructPrimitivePointersAliasField
			*out = make([]Struct_PrimitivePointers_Alias, len(*in))
			for i := range *in {
				if err := DeepCopy_wholepkg_Struct_PrimitivePointers_Alias(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.SliceSliceIntField != nil {
			in, out := &in.SliceSliceIntField, &out.SliceSliceIntField
			*out = make([][]int, len(*in))
			for i := range *in {
				if (*in)[i] != nil {
					in, out := &(*in)[i], &(*out)[i]
					*out = make([]int, len(*in))
					copy(*out, *in)
				}
			}
		}
		if in.SliceManualStructField != nil {
			in, out := &in.SliceManualStructField, &out.SliceManualStructField
			*out = make([]ManualStruct, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i].DeepCopy()
			}
		}
		if in.ManualSliceField != nil {
			out.ManualSliceField = in.ManualSliceField.DeepCopy()
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Struct_PrimitivePointers(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Struct_PrimitivePointers)
		out := out.(*Struct_Struct_PrimitivePointers)
		*out = *in
		if err := DeepCopy_wholepkg_Struct_PrimitivePointers(&in.StructField, &out.StructField, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Struct_Primitives(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Struct_Primitives)
		out := out.(*Struct_Struct_Primitives)
		*out = *in
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Struct_Slices(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Struct_Slices)
		out := out.(*Struct_Struct_Slices)
		*out = *in
		if err := DeepCopy_wholepkg_Struct_Slices(&in.StructField, &out.StructField, c); err != nil {
			return err
		}
		return nil
	}
}
