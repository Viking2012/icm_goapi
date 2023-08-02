package encoding

import (
	"bytes"
	"github.com/alexander-orban/icm_goapi/orm"
	"reflect"
	"regexp"
	"strings"
	"unicode"
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
	v := reflect.ValueOf(a)
	i := reflect.Indirect(v)
	//t := reflect.TypeOf(a)
	t := i.Type()
	var tags []string
	var indentation = strings.Repeat("\t", 1)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Implements(reflect.TypeOf((*orm.ICMEntity)(nil)).Elem()) {
			for _, tag := range extractFlatFields(a.GetFlags()) {
				tags = append(tags, tag)
			}
		} else {
			var formattedString = strings.Builder{}
			formattedString.WriteString(indentation)
			sqlTag := parseTag(field.Tag.Get("sql"))
			if sqlTag.Field == "" {
				break
			}

			if quote, _ := needsQuoting(sqlTag.Field); quote {
				formattedString.WriteString(`"`)
				formattedString.WriteString(sqlTag.Field)
				formattedString.WriteString(`"`)
			} else {
				formattedString.WriteString(sqlTag.Field)
			}
			tags = append(tags, formattedString.String())
		}
	}
	return tags
}

func extractNestedFields(a orm.ICMEntity, indentLevel int, asObject bool) []string {
	var tags []string
	if a == nil {
		return tags
	}
	v := reflect.ValueOf(a)
	indirect := reflect.Indirect(v)
	//t := reflect.TypeOf(a)
	t := indirect.Type()
	var indentation = strings.Repeat("\t", indentLevel)
	var icmEntityType = reflect.TypeOf((*orm.ICMEntity)(nil)).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		implementsEntity := field.Type.Implements(icmEntityType)
		if implementsEntity {
			var objectFieldString = strings.Builder{}
			newThing, _ := indirect.Field(i).Interface().(orm.ICMEntity)
			objectTags := extractNestedFields(newThing, indentLevel+1, true)
			objectFieldString.WriteString(indentation)
			objectFieldString.WriteString("OBJECT_CONSTRUCT(\n")
			objectFieldString.WriteString(strings.Join(objectTags, ",\n"))
			objectFieldString.WriteString(`) AS `)
			thisTag := parseTag(field.Tag.Get("sql"))
			if quote, _ := needsQuoting(thisTag.Field); quote {
				objectFieldString.WriteString(`"`)
				objectFieldString.WriteString(thisTag.Field)
				objectFieldString.WriteString(`"`)
			} else {
				objectFieldString.WriteString(thisTag.Field)
			}
			s := strings.TrimRight(objectFieldString.String(), ",\n")
			tags = append(tags, s)
		} else {
			var formattedString = strings.Builder{}
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

			if quote, _ := needsQuoting(sqlTag.Field); quote {
				formattedString.WriteString(`"`)
				formattedString.WriteString(sqlTag.Field)
				formattedString.WriteString(`"`)
			} else {
				formattedString.WriteString(sqlTag.Field)
			}

			tags = append(tags, formattedString.String())
		}
	}
	return tags
}

func needsQuoting(field string) (bool, error) {
	matched, err := regexp.Match(`^[A-Za-z_].*`, []byte(field))
	if err != nil {
		return true, err
	}
	if !matched {
		return true, nil
	}

	matched, err = regexp.Match(".*[^A-Za-z0-9_].*", []byte(field))
	if err != nil {
		return true, err
	}
	if matched {
		return true, nil
	}

	if !isUpper(field) {
		return true, nil
	}

	return false, nil
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
