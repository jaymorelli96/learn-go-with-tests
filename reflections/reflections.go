package reflections

import "reflect"

func Walk(x any, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		computeFields(val.NumField(), val.Field, fn)
	case reflect.Slice, reflect.Array:
		computeFields(val.Len(), val.Index, fn)
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			Walk(v.Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func computeFields(numberOfValues int, getField func(int) reflect.Value, fn func(string)) {
	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
