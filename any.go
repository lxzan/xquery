package parser

import "reflect"

type Any struct {
	v interface{}
}

func NewAny(v interface{}) *Any {
	var obj = new(Any)
	obj.v = v
	return obj
}

func (u *Any) Int64() int64 {
	return reflect.ValueOf(u.v).Int()
}

func (u *Any) String() string {
	return reflect.ValueOf(u.v).String()
}

func (u *Any) Interface() interface{} {
	return u.v
}
