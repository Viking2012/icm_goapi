package orm

import (
	"reflect"
)

func GetFields(a any) []string {
	t := reflect.TypeOf(a)
	var fields []string = make([]string, t.NumField())
	for i, field := range reflect.VisibleFields(t) {
		tag, found := field.Tag.Lookup("sql")
		if !found {
			tag = "not found"
		}
		fields[i] = tag
	}
	return fields
}

func NewModel(a any) map[string]int16 {
	var (
		key   string
		model map[string]int16 = make(map[string]int16)
	)

	for _, key = range GetFields(a) {
		model[key] = 0
	}
	return model
}
