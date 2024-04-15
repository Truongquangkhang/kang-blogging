package utils

import "google.golang.org/protobuf/types/known/wrapperspb"

func WrapperStringFromString(v *string) *wrapperspb.StringValue {
	if v == nil {
		return nil
	}
	return wrapperspb.String(*v)
}

func WrapperStringArrayFromStringArray(vals []*string) []*wrapperspb.StringValue {
	var wrapperVals []*wrapperspb.StringValue

	for _, val := range vals {
		wrapperVals = append(wrapperVals, WrapperStringFromString(val))
	}

	return wrapperVals
}

func WrapperInt32ArrayFromInt32Array(vals []*int32) []*wrapperspb.Int32Value {
	var wrapperVals []*wrapperspb.Int32Value

	for _, val := range vals {
		wrapperVals = append(wrapperVals, WrapperInt32FromInt32(val))
	}

	return wrapperVals
}

func WrapperInt32ArrayFromInt32ArrayPrimitive(vals []int32) []*wrapperspb.Int32Value {
	var wrapperVals []*wrapperspb.Int32Value

	for _, val := range vals {
		tmp := val
		wrapperVals = append(wrapperVals, WrapperInt32FromInt32(&tmp))
	}

	return wrapperVals
}

func WrapperInt32FromInt32(v *int32) *wrapperspb.Int32Value {
	if v == nil {
		return nil
	}
	return wrapperspb.Int32(*v)
}

func WrapperInt32FromInt64(v *int64) *wrapperspb.Int32Value {
	if v == nil {
		return nil
	}
	return wrapperspb.Int32(int32(*v))
}

func WrapperInt64FromInt64(v *int64) *wrapperspb.Int64Value {
	if v == nil {
		return nil
	}
	return wrapperspb.Int64(*v)
}

func WrapperDoubleFromDouble(v *float64) *wrapperspb.DoubleValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Double(*v)
}

func WrapperBoolFromBool(v *bool) *wrapperspb.BoolValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Bool(*v)
}

func WrapperValueString(v *wrapperspb.StringValue) *string {
	if v == nil {
		return nil
	}
	st := v.GetValue()
	return &st
}

func WrapperValuesStringArray(vals []*wrapperspb.StringValue) []*string {
	var res []*string
	for _, val := range vals {
		res = append(res, WrapperValueString(val))
	}

	return res
}

func WrapperValuesInt32Array(vals []*wrapperspb.Int32Value) []int32 {
	var res []int32
	for _, val := range vals {
		res = append(res, *WrapperValueInt32(val))
	}

	return res
}

func WrapperValuesInt64Array(vals []*wrapperspb.Int64Value) []int64 {
	var res []int64
	for _, val := range vals {
		res = append(res, *WrapperValueInt64(val))
	}

	return res
}

func WrapperValueInt32(v *wrapperspb.Int32Value) *int32 {
	if v == nil {
		return nil
	}
	st := v.GetValue()
	return &st
}

func WrapperValueInt64(v *wrapperspb.Int64Value) *int64 {
	if v == nil {
		return nil
	}
	st := v.GetValue()
	return &st
}

func WrapperValueDouble(v *wrapperspb.DoubleValue) *float64 {
	if v == nil {
		return nil
	}
	st := v.GetValue()
	return &st
}

func WrapperValueBool(v *wrapperspb.BoolValue) *bool {
	if v == nil {
		return nil
	}
	st := v.GetValue()
	return &st
}

func WrapperInt32ArrayFromInt32(v *int32) []*wrapperspb.Int32Value {
	if v == nil {
		return nil
	}
	var wrapperVals []*wrapperspb.Int32Value
	wrapperVals = append(wrapperVals, WrapperInt32FromInt32(v))
	return wrapperVals
}

func WrapperInt64ArrayFromInt64(v *int64) []*wrapperspb.Int64Value {
	if v == nil {
		return nil
	}
	var wrapperVals []*wrapperspb.Int64Value
	wrapperVals = append(wrapperVals, WrapperInt64FromInt64(v))
	return wrapperVals
}

func WrapperStringArrayFromString(v *string) []*wrapperspb.StringValue {
	if v == nil {
		return nil
	}
	var wrapperVals []*wrapperspb.StringValue
	wrapperVals = append(wrapperVals, WrapperStringFromString(v))
	return wrapperVals
}
