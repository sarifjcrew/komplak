package koplak

import (
	"reflect"
)

type Query struct{}

func (q *Query) Where(data interface{}, fn func(e interface{}) bool) (result interface{}) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(data)
		result1 := reflect.MakeSlice(val.Type(), 0, 0)
		for i := 0; i < val.Len(); i++ {
			refVal := val.Index(i)
			row := refVal.Interface()
			if row != nil {
				cek := fn(row)
				if cek {
					result1 = reflect.Append(result1, refVal)
				}
			}
		}

		result = result1.Interface()
	}

	return
}
