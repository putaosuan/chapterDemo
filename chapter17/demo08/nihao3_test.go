package demo08

import (
	"reflect"
	"testing"
)

type user8 struct {
	UserId string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *user8
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf:", st.Kind().String())
	st = st.Elem()
	t.Log("reflect.TypeOf:", st.Kind().String())
	elem = reflect.New(st) //返回一个value类型值，该值持有一个指向类型为typ的新申请的零值的指针
	t.Log("reflect.new", elem.Kind().String())
	t.Log("reflect.new.elem()", elem.Elem().Kind().String())
	model = elem.Interface().(*user8)
	elem = elem.Elem()
	elem.FieldByName("UserId").SetString("12345")
	elem.FieldByName("Name").SetString("huya")
	t.Log("model model.name", model, model.Name)
}
