package encoding

import (
	"icm_processor/orm"
	"reflect"
	"testing"
)

type testEntity struct {
	ID     string    `json:"id" sql:"ID|pk"`
	Field1 string    `json:"field-1" sql:"FIELD1"`
	Flags  testFlags `json:"flags" sql:"FLAGS"`
	Field2 string    `json:"field-2" sql:"FIELD2"`
}

func (te testEntity) GetFlags() orm.ICMEntity {
	return te.Flags
}

type testFlags struct {
	Flag1 bool `json:"flag-1" sql:"Flag: Type 1"`
	Flag2 bool `json:"flag-2" sql:"Flag: Type 2"`
	Flag3 bool `json:"flag-3" sql:"Flag: Type 3"`
}

func (tf testFlags) GetFlags() orm.ICMEntity {
	return tf
}

// TestMarshalToSelect results for cleaner tests
var nestedSelect []byte = []byte(`SELECT
	"ID",
	"FIELD1",
	OBJECT_CONSTRUCT(
		'flag-1',"Flag: Type 1",
		'flag-2',"Flag: Type 2",
		'flag-3',"Flag: Type 3") AS "FLAGS",
	"FIELD2"
FROM test`)

var flatSelect []byte = []byte(`SELECT
	"ID",
	"FIELD1",
	"Flag: Type 1",
	"Flag: Type 2",
	"Flag: Type 3",
	"FIELD2"
FROM test`)

func TestMarshalToSelect(t *testing.T) {
	type args struct {
		a        orm.ICMEntity
		database string
		flatten  bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "nested select",
			args: args{
				a:        testEntity{},
				database: "test",
				flatten:  false,
			},
			want: nestedSelect,
		},
		{
			name: "flat select",
			args: args{
				a:        testEntity{},
				database: "test",
				flatten:  true,
			},
			want: flatSelect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarshalToSelect(tt.args.a, tt.args.database, tt.args.flatten); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalToSelect(), %s\n got:\n%s\nwant:\n%s", tt.name, got, tt.want)
			}
		})
	}
}

// Test_extractFlatFields results
var standardStruct testEntity = testEntity{
	ID:     "123",
	Field1: "a string",
	Flags: testFlags{
		Flag1: false,
		Flag2: true,
		Flag3: false,
	},
	Field2: "another string",
}
var disorderedStruct testEntity = testEntity{
	Field2: "another string",
	Field1: "a string",
	ID:     "123",
	Flags: testFlags{
		Flag1: false,
		Flag2: true,
		Flag3: false,
	},
}
var flatStructFields []string = []string{`	"ID"`, `	"FIELD1"`, `	"Flag: Type 1"`, `	"Flag: Type 2"`, `	"Flag: Type 3"`, `	"FIELD2"`}

func Test_extractFlatFields(t *testing.T) {
	type args struct {
		a orm.ICMEntity
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "standard struct",
			args: args{a: standardStruct},
			want: flatStructFields,
		},
		{
			name: "disordered struct",
			args: args{a: disorderedStruct},
			want: flatStructFields,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractFlatFields(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFlatFields(), %s\n got:\n%s\nwant:\n%s", tt.name, got, tt.want)
			}
		})
	}
}

var nestedStructFields []string = []string{
	`	"ID"`,
	`	"FIELD1"`,
	`	OBJECT_CONSTRUCT(
		'flag-1',"Flag: Type 1",
		'flag-2',"Flag: Type 2",
		'flag-3',"Flag: Type 3") AS "flags"`,
	`	"FIELD2"`,
}

func Test_extractNestedFields(t *testing.T) {
	type args struct {
		a           orm.ICMEntity
		indentLevel int
		asObject    bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "standard struct",
			args: args{a: standardStruct, indentLevel: 1, asObject: false},
			want: nestedStructFields,
		},
		{
			name: "disordered struct",
			args: args{a: disorderedStruct, indentLevel: 1, asObject: false},
			want: nestedStructFields,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractNestedFields(tt.args.a, tt.args.indentLevel, tt.args.asObject); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractNestedFields(), %s\n got:\n%s\nwant:\n%s", tt.name, got, tt.want)
			}
		})
	}
}

var emptySlice []string

func Test_parseTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want sqlTag
	}{
		{
			name: "simple tag",
			args: args{tag: `simple`},
			want: sqlTag{
				Field:   "simple",
				Options: emptySlice,
			},
		},
		{
			name: "tag with option",
			args: args{tag: `simple|pk`},
			want: sqlTag{
				Field:   "simple",
				Options: []string{"pk"},
			},
		},
		{
			name: "empty tag",
			args: args{tag: ``},
			want: sqlTag{
				Field:   "",
				Options: emptySlice,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTag(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTag(), %s\n got:\n%s\nwant:\n%s", tt.name, got, tt.want)
			}
		})
	}
}
