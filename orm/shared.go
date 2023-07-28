package orm

import (
	"reflect"
)

// ICMEntity is an interface used by the sql encoder for any
// ICM entity type: bank, customer, project, or vendor
// TODO(ajo): consider changing this interface to an ICMObject
type ICMEntity interface {
	GetFlags() ICMEntity
}

func GetFields(a ICMEntity) []string {
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

func NewModel(a ICMEntity) map[string]int16 {
	var (
		key   string
		model map[string]int16 = make(map[string]int16)
	)

	for _, key = range GetFields(a) {
		model[key] = 0
	}
	return model
}
