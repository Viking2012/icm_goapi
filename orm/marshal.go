package orm

import (
	"bytes"
	"reflect"
	"strings"
)

type sqlTag struct {
	Field   string
	Options []string
}

func MarshalToSelect(a ICMEntity, database string, indent bool) []byte {
	var ret []byte
	buf := bytes.NewBufferString("SELECT\n")
	var tags []string = extractFields(a, 1, false)

	s := strings.Join(tags, ",\n")

	buf.WriteString(strings.TrimRight(s, ",\n\t"))
	buf.WriteString("\n")
	buf.WriteString("FROM ")
	buf.WriteString(database)

	ret = buf.Bytes()
	if !indent {
		return bytes.ReplaceAll(bytes.ReplaceAll(ret, []byte("\n"), []byte("")), []byte("\t"), []byte(""))
	} else {
		return ret
	}
}

func extractFields(a ICMEntity, indentLevel int, asObject bool) []string {
	t := reflect.TypeOf(a)
	var tags []string
	var indentation = strings.Repeat("\t", indentLevel)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Implements(reflect.TypeOf((*ICMEntity)(nil)).Elem()) {
			//tags = append(tags, indentation + "OBJECT_CONSTRUCT(")
			objectTags := extractFields(a.GetFlags(), indentLevel+1, true)
			s := indentation + "OBJECT_CONSTRUCT(\n" + strings.Join(objectTags, ",\n") + ") AS " + field.Tag.Get("json")
			tags = append(tags, strings.TrimRight(s, ",\n"))
			//for _, f := range extractFields(a.GetFlags(),indentLevel+1, true) {
			//	tags = append(tags, indentation + f)
			//}
			//tags = append(tags, indentation + ") AS " + field.Tag.Get("json"))
		} else {
			var formattedString strings.Builder = strings.Builder{}
			formattedString.WriteString(indentation)
			sqlTag := parseTag(field.Tag.Get("sql"))
			if sqlTag.Field == "" {
				break
			}
			json := parseTag(field.Tag.Get("json"))

			if asObject {
				formattedString.WriteString("'")
				formattedString.WriteString(json.Field)
				formattedString.WriteString("',")
			}

			formattedString.WriteString(`"`)
			formattedString.WriteString(sqlTag.Field)
			formattedString.WriteString(`"`)
			tags = append(tags, formattedString.String())
		}
	}
	//tags[len(tags) - 1] = strings.TrimRight(tags[len(tags) - 1], ",\n\t")
	return tags
}

func parseTag(tag string) sqlTag {
	var (
		field   string
		opts    string
		options []string
	)
	field, opts, _ = strings.Cut(tag, "|")
	if opts != "" {
		options = strings.Split(opts, ",")
	}
	return sqlTag{
		Field:   field,
		Options: options,
	}
}
