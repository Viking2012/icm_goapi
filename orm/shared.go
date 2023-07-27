package orm

import (
	"reflect"
)

// // BoolFromFloat from https://stackoverflow.com/a/37214476
// type BoolFromFloat bool
//
//	func (bit *BoolFromFloat) UnmarshalJSON(data []byte) error {
//		asString := string(data)
//		if asString == "1" || asString == "true" {
//			*bit = true
//		} else if asString == "0" || asString == "false" {
//			*bit = false
//		} else {
//			return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
//		}
//		return nil
//	}
//
//	func (cb *BoolFromFloat) Scan(src interface{}) error {
//		var scan BoolFromFloat
//		switch src.(type) {
//		case float32:
//			switch src {
//			case 0:
//				scan = false
//			case 1:
//				scan = true
//			default:
//				return errors.New("incompatible int value for a IntFromFloat")
//			}
//		case float64:
//			switch src {
//			case 0:
//				scan = false
//			case 1:
//				scan = true
//			default:
//				return errors.New("incompatible int value for a IntFromFloat")
//			}
//		default:
//			return errors.New("incompatible value type for a IntFromFloat")
//		}
//
//		*cb = scan
//		return nil
//	}
//
//	func (cb *BoolFromFloat) Value() (driver.Value, error) {
//		if *cb {
//			return driver.Value(float64(1)), nil
//		} else {
//			return driver.Value(float64(0)), nil
//		}
//	}
//
// type IntFromFloat int64
//
//	func (ci *IntFromFloat) Scan(src interface{}) error {
//		s := fmt.Sprintf("%.0f", src)
//		i, err := strconv.Atoi(s)
//		if err != nil {
//			return err
//		}
//		*ci = IntFromFloat(int64(i))
//		return err
//	}
//
//	func (ci *IntFromFloat) Value() (driver.Value, error) {
//		return driver.Value(float64(*ci)), nil
//	}
//
//	func (ci *IntFromFloat) UnmarshalJSON(data []byte) error {
//		f, err := strconv.ParseFloat(string(data), 64)
//		if err != nil {
//			return err
//		}
//		s := fmt.Sprintf("%.0f", f)
//		i, err := strconv.Atoi(s)
//		if err != nil {
//			return err
//		}
//		*ci = IntFromFloat(int64(i))
//		return err
//	}
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
