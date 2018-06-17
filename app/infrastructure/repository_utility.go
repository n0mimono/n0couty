package infrastructure

import (
	"reflect"
)

func Copy(ps interface{}, pd interface{}, td interface{}) {
	t := reflect.TypeOf(td).Elem()

	vs := reflect.ValueOf(ps).Elem()
	vd := reflect.ValueOf(pd).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fs := vs.FieldByName(field.Name)
		fd := vd.FieldByName(field.Name)
		fs.Set(fd)
	}
}

func CopyFrom(ps interface{}, pd interface{}) {
	Copy(ps, pd, pd)
}

func CopyTo(ps interface{}, pd interface{}) {
	Copy(pd, ps, pd)
}
