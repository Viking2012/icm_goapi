package orm

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// ICMEntity is an interface used by the sql encoder for any
// ICM entity type: bank, customer, project, or vendor
// TODO(ajo): consider changing this interface to an ICMObject
type ICMEntity interface {
	IsICMEntity() bool
	//LoadFromRow(row *sql.Rows) error
}

func GetFields(a *ICMEntity) []string {
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

func NewModel(a *ICMEntity) map[string]int16 {
	var (
		key   string
		model map[string]int16 = make(map[string]int16)
	)

	for _, key = range GetFields(a) {
		model[key] = 0
	}
	return model
}

// BoolFromFloat from https://stackoverflow.com/a/37214476
type BoolFromFloat bool

func (bit *BoolFromFloat) UnmarshalJSON(data []byte) error {
	var scan BoolFromFloat
	f, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return errors.New(fmt.Sprintf("incompatible float value for a BoolFromFloat UnarshalJSON: %s", data))
	}
	s := fmt.Sprintf("%.0f", f)
	switch s {
	case "0":
		scan = false
	case "1":
		scan = true
	default:
		return errors.New(fmt.Sprintf("incompatible float value for a BoolFromFloat UnarshalJSON: %s", s))
	}
	*bit = scan
	return nil
}

func (bit *BoolFromFloat) Scan(src interface{}) error {
	var scan BoolFromFloat
	f, ok := src.(float64)
	if !ok {
		return errors.New(fmt.Sprintf("incompatible float value for a BoolFromFloat Scan Float: %s", src))
	}
	s := fmt.Sprintf("%.0f", f)
	switch s {
	case "0":
		scan = false
	case "1":
		scan = true
	default:
		return errors.New(fmt.Sprintf("incompatible float value for a BoolFromFloat Scan: %s", src))
	}
	*bit = scan
	return nil
}

func (bit *BoolFromFloat) Value() (driver.Value, error) {
	if *bit {
		return driver.Value(float64(1)), nil
	} else {
		return driver.Value(float64(0)), nil
	}
}

type IntFromFloat int64

func (ci *IntFromFloat) Scan(src interface{}) error {
	s := fmt.Sprintf("%.0f", src)
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*ci = IntFromFloat(int64(i))
	return err
}

func (ci *IntFromFloat) Value() (driver.Value, error) {
	return driver.Value(float64(*ci)), nil
}

func (ci *IntFromFloat) UnmarshalJSON(data []byte) error {
	f, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	s := fmt.Sprintf("%.0f", f)
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*ci = IntFromFloat(int64(i))
	return err
}

var nullTime = time.Time{}

type NullTime struct {
	time.Time
}

func (t *NullTime) Value() (driver.Value, error) {
	if t.Time == nullTime {
		return nil, nil
	}
	return t.Time, nil
}

func (t *NullTime) Scan(src interface{}) error {
	if src == nil {
		t.Time = nullTime
		return nil
	}
	scannedTime, ok := src.(time.Time)
	if !ok {
		return errors.New(fmt.Sprintf("incompatible time value for a NullTime Scan: %s", src))
	}
	t.Time = scannedTime
	return nil
}
