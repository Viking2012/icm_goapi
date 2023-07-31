package encoding

import (
	"bytes"
	"icm_processor/orm"
	"reflect"
	"strings"
)

type sqlTag struct {
	Field   string
	Options []string
}

func (s *sqlTag) ContainsOption(tag string) bool {
	for i := range s.Options {
		if tag == s.Options[i] {
			return true
		}
	}
	return false
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

// MarshalToSelect return a byte array which represent the equivalent SELECT statement
// which would return a representation of an internal struct (such as Vendor, Bank, etc.)
// The flatten argument determines if the SELECT statement will either
//   - true:  returns a simple SELECT statement, which calls for each field, nested or not,
//     to be returned as an individual column - roughly equivalent to a `SELECT *` statement
//   - or -
//   - false: returns a query which takes struct fields which are, themselves,
//     internal structs (like Flags structs) and converts them into a single return field
//     through and OBJECT_CONSTRUCT() command
//
// sql fields names are taken from the `sql` tags on the struct
// field names for the generated OBJECT are taken from the `json` tags of the struct
func MarshalToSelect(a orm.ICMEntity, database string, flatten bool) []byte {
	buf := bytes.NewBufferString("SELECT\n")
	var tags []string
	if flatten {
		tags = extractFlatFields(a)
	} else {
		tags = extractNestedFields(a, 1, false)

	}

	s := strings.Join(tags, ",\n")

	buf.WriteString(strings.TrimRight(s, ",\n\t"))
	buf.WriteString("\n")
	buf.WriteString("FROM ")
	buf.WriteString(database)

	return buf.Bytes()
}

func extractFlatFields(a orm.ICMEntity) []string {
	t := reflect.TypeOf(a)
	var tags []string
	var indentation = strings.Repeat("\t", 1)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Implements(reflect.TypeOf((*orm.ICMEntity)(nil)).Elem()) {
			for _, tag := range extractFlatFields(a.GetFlags()) {
				tags = append(tags, tag)
			}
		} else {
			var formattedString strings.Builder = strings.Builder{}
			formattedString.WriteString(indentation)
			sqlTag := parseTag(field.Tag.Get("sql"))
			if sqlTag.Field == "" {
				break
			}

			formattedString.WriteString(`"`)
			formattedString.WriteString(sqlTag.Field)
			formattedString.WriteString(`"`)
			tags = append(tags, formattedString.String())
		}
	}
	return tags
}

func extractNestedFields(a orm.ICMEntity, indentLevel int, asObject bool) []string {
	t := reflect.TypeOf(a)
	var tags []string
	var indentation = strings.Repeat("\t", indentLevel)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Implements(reflect.TypeOf((*orm.ICMEntity)(nil)).Elem()) {
			var objectFieldSring strings.Builder = strings.Builder{}
			objectTags := extractNestedFields(a.GetFlags(), indentLevel+1, true)
			objectFieldSring.WriteString(indentation)
			objectFieldSring.WriteString("OBJECT_CONSTRUCT(\n")
			objectFieldSring.WriteString(strings.Join(objectTags, ",\n"))
			objectFieldSring.WriteString(`) AS "`)
			objectFieldSring.WriteString(field.Tag.Get("sql"))
			objectFieldSring.WriteString(`"`)
			s := strings.TrimRight(objectFieldSring.String(), ",\n")
			tags = append(tags, s)
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
	return tags
}
